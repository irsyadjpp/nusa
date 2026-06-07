package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/middleware"
	authHandler "github.com/nusa/backend/modules/auth"
	curriculumHandler "github.com/nusa/backend/modules/curriculum"
	learningPlanningHandler "github.com/nusa/backend/modules/learning_planning"
	assessmentHandler "github.com/nusa/backend/modules/assessment"
	reportingHandler "github.com/nusa/backend/modules/reporting"
	roleHandler "github.com/nusa/backend/modules/roles"
	schoolHandler "github.com/nusa/backend/modules/schools"
	userHandler "github.com/nusa/backend/modules/users"
	jwtService "github.com/nusa/backend/pkg/jwt"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter(
	authHandler *authHandler.Handler,
	userHandler *userHandler.Handler,
	schoolHandler *schoolHandler.Handler,
	roleHandler *roleHandler.Handler,
	curriculumHandler *curriculumHandler.Handler,
	learningPlanningHandler *learningPlanningHandler.Handler,
	assessmentHandler *assessmentHandler.Handler,
	reportingHandler *reportingHandler.Handler,
	jwtService *jwtService.Service,
	userRepo interface{},
	schoolRepo interface{},
) *Router {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	r := &Router{engine: engine}

	r.setupRoutes(authHandler, userHandler, schoolHandler, roleHandler, curriculumHandler, learningPlanningHandler, assessmentHandler, reportingHandler, jwtService, userRepo, schoolRepo)

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
	reportingHandler *reportingHandler.Handler,
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
		{
			curriculum.POST("/subjects", curriculumHandler.CreateCurriculumSubject)
			curriculum.GET("/subjects", curriculumHandler.ListCurriculumSubjects)
			curriculum.GET("/subjects/:id", curriculumHandler.GetCurriculumSubject)
			
			curriculum.POST("/cp/import", curriculumHandler.ImportCP)
			curriculum.GET("/cp", curriculumHandler.ListCPs)
			curriculum.GET("/cp/:id", curriculumHandler.GetCP)
		}

		learningPlanning := protected.Group("/learning-planning")
		{
			learningPlanning.POST("/tp-sets", learningPlanningHandler.CreateTPSet)
			learningPlanning.GET("/tp-sets", learningPlanningHandler.ListTPSets)
			learningPlanning.GET("/tp-sets/:id", learningPlanningHandler.GetTPSet)
			learningPlanning.POST("/tp-sets/:id/approve", learningPlanningHandler.ApproveTPSet)
			
			learningPlanning.POST("/atp-sets", learningPlanningHandler.CreateATPSet)
			learningPlanning.GET("/atp-sets", learningPlanningHandler.ListATPSets)
			
			learningPlanning.POST("/modul-ajar-sets", learningPlanningHandler.CreateModulAjarSet)
			learningPlanning.GET("/modul-ajar-sets", learningPlanningHandler.ListModulAjarSets)
		}

		assessment := protected.Group("/assessment")
		{
			assessment.POST("", assessmentHandler.CreateAssessment)
			assessment.GET("", assessmentHandler.ListAssessments)
			assessment.GET("/:id", assessmentHandler.GetAssessment)
			
			assessment.POST("/rubrics", assessmentHandler.CreateRubric)
			assessment.GET("/rubrics", assessmentHandler.ListRubrics)
			
			assessment.POST("/evidences", assessmentHandler.CreateEvidence)
			assessment.GET("/evidences", assessmentHandler.ListEvidences)
			
			assessment.POST("/evaluations", assessmentHandler.CreateEvaluation)
			assessment.GET("/evaluations", assessmentHandler.ListEvaluations)
		}

		reporting := protected.Group("/reporting")
		{
			reporting.POST("/narrative-reports", reportingHandler.CreateNarrativeReport)
			reporting.GET("/narrative-reports", reportingHandler.ListNarrativeReports)
			reporting.GET("/narrative-reports/:id", reportingHandler.GetNarrativeReport)
		}
	}
}

func (r *Router) GetEngine() *gin.Engine {
	return r.engine
}
