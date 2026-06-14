package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/handler"
	"github.com/nusa/backend/internal/middleware"
	achievementModule "github.com/nusa/backend/modules/achievement"
	assessmentHandler "github.com/nusa/backend/modules/assessment"
	authHandler "github.com/nusa/backend/modules/auth"
	curriculumHandler "github.com/nusa/backend/modules/curriculum"
	learningPlanningHandler "github.com/nusa/backend/modules/learning_planning"
	reportingHandler "github.com/nusa/backend/modules/reporting"
	roleHandler "github.com/nusa/backend/modules/roles"
	schoolHandler "github.com/nusa/backend/modules/schools"
	userHandler "github.com/nusa/backend/modules/users"
	jwtService "github.com/nusa/backend/pkg/jwt"
)

type Router struct {
	engine *gin.Engine
}

func (r *Router) GetEngine() *gin.Engine {
	return r.engine
}

func NewRouter(
	authHandler *authHandler.Handler,
	userHandler *userHandler.Handler,
	schoolHandler *schoolHandler.Handler,
	roleHandler *roleHandler.Handler,
	curriculumHandler *curriculumHandler.Handler,
	learningPlanningHandler *learningPlanningHandler.Handler,
	assessmentHandler *assessmentHandler.Handler,
	achievementHandler *achievementModule.Handler,
	reportingHandler *reportingHandler.Handler,
	tpSetHandler *handler.TPSetHandler,
	academicYearHandler *handler.AcademicYearHandler,
	semesterHandler *handler.SemesterHandler,
	curriculumGovernanceHandler *handler.CurriculumGovernanceHandler,
	systemConfigurationHandler *handler.SystemConfigurationHandler,
	jwtService *jwtService.Service,
	userRepo interface{},
	schoolRepo interface{},
) *Router {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	r := &Router{engine: engine}

	r.setupRoutes(authHandler, userHandler, schoolHandler, roleHandler, curriculumHandler, learningPlanningHandler, assessmentHandler, achievementHandler, reportingHandler, tpSetHandler, academicYearHandler, semesterHandler, curriculumGovernanceHandler, systemConfigurationHandler, jwtService, userRepo, schoolRepo)

	return r
}

func (r *Router) setupRoutes(
	authHandler *authHandler.Handler,
	userHandler *userHandler.Handler,
	schoolHandler *schoolHandler.Handler,
	roleHandler *roleHandler.Handler,
	curriculumHandler *curriculumHandler.Handler,
	learningPlanningHandler *learningPlanningHandler.Handler,
	assessmentHandler *assessmentHandler.Handler,
	achievementHandler *achievementModule.Handler,
	reportingHandler *reportingHandler.Handler,
	tpSetHandler *handler.TPSetHandler,
	academicYearHandler *handler.AcademicYearHandler,
	semesterHandler *handler.SemesterHandler,
	curriculumGovernanceHandler *handler.CurriculumGovernanceHandler,
	systemConfigurationHandler *handler.SystemConfigurationHandler,
	jwtService *jwtService.Service,
	userRepo interface{},
	schoolRepo interface{},
) {
	r.engine.Use(middleware.Recovery())
	r.engine.Use(middleware.CORS())
	r.engine.Use(middleware.RequestID())

	r.engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})
	r.engine.GET("/ready", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ready"})
	})
	r.engine.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": "1.0.0", "name": "NUSA Backend API"})
	})

	// Scalar API Documentation
	scalarHandler := handler.NewScalarHandler()
	r.engine.GET("/scalar", scalarHandler.ServeScalar)
	r.engine.GET("/swagger", scalarHandler.ServeSwaggerUI)
	r.engine.GET("/openapi.json", scalarHandler.ServeOpenAPISpec)

	// Handle favicon request - return 204 (no content) to avoid 404 logs
	r.engine.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	public := r.engine.Group("/api/v1/public")
	{
		auth := public.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.Refresh)
		}
	}

	protected := r.engine.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware(jwtService))
	{
		auth := protected.Group("/auth")
		{
			auth.POST("/logout", authHandler.Logout)
			auth.GET("/me", authHandler.Me)
		}

		users := protected.Group("/users")
		{
			users.POST("", middleware.RequirePermission("user:CREATE"), userHandler.CreateUser)
			users.GET("", middleware.RequirePermission("user:READ"), userHandler.GetUsers)
			users.GET("/:id", middleware.RequirePermission("user:READ"), userHandler.GetUser)
			users.PUT("/:id", middleware.RequirePermission("user:UPDATE"), userHandler.UpdateUser)
			users.PATCH("/:id/status", middleware.RequirePermission("user:UPDATE"), userHandler.UpdateUserStatus)
		}

		schools := protected.Group("/schools")
		{
			schools.POST("", middleware.RequirePermission("school:CREATE"), schoolHandler.CreateSchool)
			schools.GET("", middleware.RequirePermission("school:READ"), schoolHandler.GetSchools)
			schools.GET("/:id", middleware.RequirePermission("school:READ"), schoolHandler.GetSchool)
			schools.PUT("/:id", middleware.RequirePermission("school:UPDATE"), schoolHandler.UpdateSchool)
			schools.PATCH("/:id/status", middleware.RequirePermission("school:UPDATE"), schoolHandler.UpdateSchoolStatus)
		}

		roles := protected.Group("/roles")
		{
			roles.GET("", middleware.RequirePermission("user:READ"), roleHandler.GetRoles)
			roles.GET("/:id", middleware.RequirePermission("user:READ"), roleHandler.GetRole)
			roles.GET("/:id/permissions", middleware.RequirePermission("user:READ"), roleHandler.GetPermissions)
			roles.POST("", middleware.RequireRole(domain.RoleSystemAdmin), middleware.RequirePermission("user:CREATE"), roleHandler.CreateRole)
			roles.DELETE("/:id", middleware.RequireRole(domain.RoleSystemAdmin), middleware.RequirePermission("user:DELETE"), roleHandler.DeleteRole)
			roles.PUT("/:id", middleware.RequireRole(domain.RoleSystemAdmin), middleware.RequirePermission("user:UPDATE"), roleHandler.UpdateRole)
			roles.POST("/:id/permissions", middleware.RequirePermission("user:UPDATE"), roleHandler.AddPermission)
			roles.DELETE("/:id/permissions", middleware.RequirePermission("user:UPDATE"), roleHandler.RemovePermission)
		}

		// Education Domain Routes
		curriculum := protected.Group("/curriculum")
		curriculum.Use(middleware.ReadOnlyMiddleware(middleware.RoleSystemAdmin, middleware.RoleSchoolAdmin, middleware.RoleTeacher))
		{
			curriculum.POST("/subjects", curriculumHandler.CreateCurriculumSubject)
			curriculum.GET("/subjects", curriculumHandler.ListCurriculumSubjects)
			curriculum.GET("/subjects/:id", curriculumHandler.GetCurriculumSubject)
			curriculum.PUT("/subjects/:id", curriculumHandler.UpdateCurriculumSubject)
			curriculum.DELETE("/subjects/:id", curriculumHandler.DeleteCurriculumSubject)
			curriculum.POST("/phases", curriculumHandler.CreateCurriculumPhase)
			curriculum.GET("/phases", curriculumHandler.ListCurriculumPhases)
			curriculum.GET("/phases/:id", curriculumHandler.GetCurriculumPhase)
			curriculum.PUT("/phases/:id", curriculumHandler.UpdateCurriculumPhase)
			curriculum.DELETE("/phases/:id", curriculumHandler.DeleteCurriculumPhase)
			curriculum.POST("/elements", curriculumHandler.CreateCurriculumElement)
			curriculum.GET("/elements", curriculumHandler.ListCurriculumElements)
			curriculum.GET("/elements/:id", curriculumHandler.GetCurriculumElement)
			curriculum.PUT("/elements/:id", curriculumHandler.UpdateCurriculumElement)
			curriculum.DELETE("/elements/:id", curriculumHandler.DeleteCurriculumElement)
			curriculum.POST("/subelements", curriculumHandler.CreateCurriculumSubelement)
			curriculum.GET("/subelements", curriculumHandler.ListCurriculumSubelements)
			curriculum.GET("/subelements/:id", curriculumHandler.GetCurriculumSubelement)
			curriculum.PUT("/subelements/:id", curriculumHandler.UpdateCurriculumSubelement)
			curriculum.DELETE("/subelements/:id", curriculumHandler.DeleteCurriculumSubelement)
			curriculum.POST("/cp", curriculumHandler.CreateCP)
			curriculum.PUT("/cp/:id", curriculumHandler.UpdateCP)
			curriculum.DELETE("/cp/:id", curriculumHandler.DeleteCP)
			curriculum.GET("/cp/export", curriculumHandler.ExportCPs)

			curriculum.POST("/cp/import", curriculumHandler.ImportCP)
			curriculum.GET("/cp", curriculumHandler.ListCPs)
			curriculum.GET("/cp/:id", curriculumHandler.GetCP)
		}

		learningPlanning := protected.Group("/learning-planning")
		learningPlanning.Use(middleware.RequirePermission("tp:READ"))
		{
			// TP Set routes
			learningPlanning.POST("/tp-sets", middleware.RequirePermission("tp:CREATE"), learningPlanningHandler.CreateTPSet)
			learningPlanning.GET("/tp-sets", learningPlanningHandler.ListTPSets)
			learningPlanning.GET("/tp-sets/:id", learningPlanningHandler.GetTPSet)
			learningPlanning.POST("/tp-sets/:id/approve", middleware.RequirePermission("tp:APPROVE"), learningPlanningHandler.ApproveTPSet)
			learningPlanning.PUT("/tp-sets/:id", middleware.RequirePermission("tp:UPDATE"), learningPlanningHandler.UpdateTPSet)
			learningPlanning.GET("/tp-sets/:id/versions", learningPlanningHandler.GetTPSetVersions)

			// Individual TP routes
			learningPlanning.POST("/tps", middleware.RequirePermission("tp:CREATE"), learningPlanningHandler.CreateTP)
			learningPlanning.GET("/tps", learningPlanningHandler.ListTPs)
			learningPlanning.GET("/tps/:id", learningPlanningHandler.GetTP)

			// ATP Set routes
			learningPlanning.POST("/atp-sets", middleware.RequirePermission("tp:CREATE"), learningPlanningHandler.CreateATPSet)
			learningPlanning.GET("/atp-sets", learningPlanningHandler.ListATPSets)
			learningPlanning.GET("/atp-sets/:id", learningPlanningHandler.GetATPSet)
			learningPlanning.PUT("/atp-sets/:id", middleware.RequirePermission("tp:UPDATE"), learningPlanningHandler.UpdateATPSet)
			learningPlanning.DELETE("/atp-sets/:id", middleware.RequirePermission("tp:DELETE"), learningPlanningHandler.DeleteATPSet)
			learningPlanning.POST("/atp-sets/:id/approve", middleware.RequirePermission("tp:APPROVE"), learningPlanningHandler.ApproveATPSet)

			// Individual ATP routes
			learningPlanning.POST("/atps", middleware.RequirePermission("tp:CREATE"), learningPlanningHandler.CreateATP)
			learningPlanning.GET("/atps", learningPlanningHandler.ListATPs)
			learningPlanning.GET("/atps/:id", learningPlanningHandler.GetATP)
			learningPlanning.PUT("/atps/:id", middleware.RequirePermission("tp:UPDATE"), learningPlanningHandler.UpdateATP)
			learningPlanning.DELETE("/atps/:id", middleware.RequirePermission("tp:DELETE"), learningPlanningHandler.DeleteATP)

			// Modul Ajar Set routes
			learningPlanning.POST("/modul-ajar-sets", middleware.RequirePermission("tp:CREATE"), learningPlanningHandler.CreateModulAjarSet)
			learningPlanning.GET("/modul-ajar-sets", learningPlanningHandler.ListModulAjarSets)
			learningPlanning.GET("/modul-ajar-sets/:id", learningPlanningHandler.GetModulAjarSet)
			learningPlanning.PUT("/modul-ajar-sets/:id", middleware.RequirePermission("tp:UPDATE"), learningPlanningHandler.UpdateModulAjarSet)
			learningPlanning.DELETE("/modul-ajar-sets/:id", middleware.RequirePermission("tp:DELETE"), learningPlanningHandler.DeleteModulAjarSet)
			learningPlanning.POST("/modul-ajar-sets/:id/approve", middleware.RequirePermission("tp:APPROVE"), learningPlanningHandler.ApproveModulAjarSet)

			// Individual Modul Ajar routes
			learningPlanning.POST("/modul-ajar", middleware.RequirePermission("tp:CREATE"), learningPlanningHandler.CreateModulAjar)
			learningPlanning.GET("/modul-ajar", learningPlanningHandler.ListModulAjars)
			learningPlanning.GET("/modul-ajar/:id", learningPlanningHandler.GetModulAjar)
			learningPlanning.PUT("/modul-ajar/:id", middleware.RequirePermission("tp:UPDATE"), learningPlanningHandler.UpdateModulAjar)
			learningPlanning.DELETE("/modul-ajar/:id", middleware.RequirePermission("tp:DELETE"), learningPlanningHandler.DeleteModulAjar)
		}

		assessment := protected.Group("/assessment")
		assessment.Use(middleware.RequirePermission("assessment:READ"))
		{
			assessment.POST("", middleware.RequirePermission("assessment:CREATE"), assessmentHandler.CreateAssessment)
			assessment.GET("", assessmentHandler.ListAssessments)
			assessment.GET("/:id", assessmentHandler.GetAssessment)
			assessment.PUT("/:id", middleware.RequirePermission("assessment:UPDATE"), assessmentHandler.UpdateAssessment)
			assessment.POST("/:id/approve", middleware.RequirePermission("assessment:APPROVE"), assessmentHandler.ApproveAssessment)

			assessment.POST("/rubrics", middleware.RequirePermission("assessment:CREATE"), assessmentHandler.CreateRubric)
			assessment.GET("/rubrics", assessmentHandler.ListRubrics)
			assessment.GET("/rubrics/:id", assessmentHandler.GetRubric)
			assessment.PUT("/rubrics/:id", middleware.RequirePermission("assessment:UPDATE"), assessmentHandler.UpdateRubric)
			assessment.DELETE("/rubrics/:id", middleware.RequirePermission("assessment:DELETE"), assessmentHandler.DeleteRubric)

			assessment.POST("/evidences/upload", middleware.RequirePermission("assessment:CREATE"), assessmentHandler.UploadEvidence)
			assessment.GET("/evidences/:id", assessmentHandler.GetEvidence)
			assessment.POST("/evidences", middleware.RequirePermission("assessment:CREATE"), assessmentHandler.CreateEvidence)
			assessment.GET("/evidences", assessmentHandler.ListEvidences)
			assessment.PUT("/evidences/:id", middleware.RequirePermission("assessment:UPDATE"), assessmentHandler.UpdateEvidence)
			assessment.DELETE("/evidences/:id", middleware.RequirePermission("assessment:DELETE"), assessmentHandler.DeleteEvidence)

			assessment.POST("/evaluations", middleware.RequirePermission("assessment:CREATE"), assessmentHandler.CreateEvaluation)
			assessment.GET("/evaluations", assessmentHandler.ListEvaluations)
			assessment.GET("/evaluations/:id", assessmentHandler.GetEvaluation)
			assessment.PUT("/evaluations/:id", middleware.RequirePermission("assessment:UPDATE"), assessmentHandler.UpdateEvaluation)
			assessment.GET("/evaluations/history/:evidence_id", assessmentHandler.GetEvaluationHistory)
			assessment.GET("/evaluations/:id/feedback-history", assessmentHandler.GetEvaluationFeedbackHistory)
		}

		reporting := protected.Group("/reporting")
		reporting.Use(middleware.RequirePermission("reporting:READ"))
		{
			reporting.POST("/narrative-reports", middleware.RequirePermission("reporting:CREATE"), reportingHandler.CreateNarrativeReport)
			reporting.GET("/narrative-reports", reportingHandler.ListNarrativeReports)
			reporting.GET("/narrative-reports/:id", reportingHandler.GetNarrativeReport)
			reporting.PUT("/narrative-reports/:id", middleware.RequirePermission("reporting:UPDATE"), reportingHandler.UpdateNarrativeReport)
			reporting.DELETE("/narrative-reports/:id", middleware.RequirePermission("reporting:DELETE"), reportingHandler.DeleteNarrativeReport)
			reporting.POST("/narrative-reports/:id/refresh-achievement", middleware.RequirePermission("reporting:UPDATE"), reportingHandler.RefreshReportAchievement)
		}

		// Achievement Routes
		students := protected.Group("/students")
		students.Use(middleware.RequirePermission("reporting:READ"))
		{
			students.GET("/:id/achievement", achievementHandler.GetStudentAchievement)
			students.GET("/:id/progress", achievementHandler.GetStudentProgress)
		}

		classes := protected.Group("/classes")
		classes.Use(middleware.RequirePermission("reporting:READ"))
		{
			classes.GET("/:id/achievement", achievementHandler.GetClassAchievement)
		}

		// Sprint 4: Academic Foundation Routes
		academicYears := protected.Group("/academic-years")
		academicYears.Use(middleware.RequirePermission("academic_year:READ"))
		{
			academicYears.POST("", middleware.RequirePermission("academic_year:CREATE"), academicYearHandler.CreateAcademicYear)
			academicYears.GET("", academicYearHandler.ListAcademicYears)
			academicYears.GET("/:id", academicYearHandler.GetAcademicYear)
			academicYears.PUT("/:id", middleware.RequirePermission("academic_year:UPDATE"), academicYearHandler.UpdateAcademicYear)
			academicYears.POST("/:id/activate", middleware.RequirePermission("academic_year:ACTIVATE"), academicYearHandler.ActivateAcademicYear)
			academicYears.POST("/:id/archive", middleware.RequirePermission("academic_year:ARCHIVE"), academicYearHandler.ArchiveAcademicYear)
		}

		semesters := protected.Group("/semesters")
		semesters.Use(middleware.RequirePermission("semester:READ"))
		{
			semesters.POST("", middleware.RequirePermission("semester:CREATE"), semesterHandler.CreateSemester)
			semesters.GET("", semesterHandler.ListSemesters)
			semesters.GET("/:id", semesterHandler.GetSemester)
			semesters.PUT("/:id", middleware.RequirePermission("semester:UPDATE"), semesterHandler.UpdateSemester)
			semesters.DELETE("/:id", middleware.RequirePermission("semester:DELETE"), semesterHandler.DeleteSemester)
		}

		subjectCategories := protected.Group("/subject-categories")
		subjectCategories.Use(middleware.RequirePermission("subject_category:READ"))
		{
			subjectCategories.POST("", middleware.RequirePermission("subject_category:CREATE"), curriculumGovernanceHandler.CreateSubjectCategory)
			subjectCategories.GET("", curriculumGovernanceHandler.ListSubjectCategories)
			subjectCategories.PUT("/:id", middleware.RequirePermission("subject_category:UPDATE"), curriculumGovernanceHandler.UpdateSubjectCategory)
			subjectCategories.DELETE("/:id", middleware.RequirePermission("subject_category:DELETE"), curriculumGovernanceHandler.DeleteSubjectCategory)
		}

		graduateProfileDimensions := protected.Group("/graduate-profile-dimensions")
		graduateProfileDimensions.Use(middleware.RequirePermission("graduate_profile:READ"))
		{
			graduateProfileDimensions.POST("", middleware.RequirePermission("graduate_profile:CREATE"), curriculumGovernanceHandler.CreateGraduateProfileDimension)
			graduateProfileDimensions.GET("", curriculumGovernanceHandler.ListGraduateProfileDimensions)
			graduateProfileDimensions.PUT("/:id", middleware.RequirePermission("graduate_profile:UPDATE"), curriculumGovernanceHandler.UpdateGraduateProfileDimension)
			graduateProfileDimensions.DELETE("/:id", middleware.RequirePermission("graduate_profile:DELETE"), curriculumGovernanceHandler.DeleteGraduateProfileDimension)
		}

		cpAlignments := protected.Group("/cp-alignments")
		cpAlignments.Use(middleware.RequirePermission("cp_alignment:READ"))
		{
			cpAlignments.POST("", middleware.RequirePermission("cp_alignment:CREATE"), curriculumGovernanceHandler.CreateCPAlignment)
			cpAlignments.POST("/bulk", middleware.RequirePermission("cp_alignment:CREATE"), curriculumGovernanceHandler.CreateCPAlignmentBulk)
			cpAlignments.GET("", curriculumGovernanceHandler.ListCPAlignments)
			cpAlignments.GET("/report", curriculumGovernanceHandler.GenerateCPAlignmentReport)
			cpAlignments.PUT("/:id", middleware.RequirePermission("cp_alignment:UPDATE"), curriculumGovernanceHandler.UpdateCPAlignment)
			cpAlignments.DELETE("/:id", middleware.RequirePermission("cp_alignment:DELETE"), curriculumGovernanceHandler.DeleteCPAlignment)
		}

		systemConfigurations := protected.Group("/system-configurations")
		systemConfigurations.Use(middleware.RequirePermission("system_config:READ"))
		{
			systemConfigurations.POST("", middleware.RequirePermission("system_config:CREATE"), systemConfigurationHandler.CreateSystemConfiguration)
			systemConfigurations.GET("", systemConfigurationHandler.ListSystemConfigurations)
			systemConfigurations.GET("/:id", systemConfigurationHandler.GetSystemConfiguration)
			systemConfigurations.GET("/by-key/:key", systemConfigurationHandler.GetSystemConfigurationByKey)
			systemConfigurations.PUT("/:id", middleware.RequirePermission("system_config:UPDATE"), systemConfigurationHandler.UpdateSystemConfiguration)
			systemConfigurations.DELETE("/:id", middleware.RequirePermission("system_config:DELETE"), systemConfigurationHandler.DeleteSystemConfiguration)
		}
	}
}
