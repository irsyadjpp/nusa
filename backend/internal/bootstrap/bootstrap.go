package bootstrap

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nusa/backend/internal/config"
	"github.com/nusa/backend/internal/db"
	"github.com/nusa/backend/internal/logger"
	"github.com/nusa/backend/internal/middleware"
	"github.com/nusa/backend/internal/repository"
	"github.com/nusa/backend/internal/router"
	"github.com/nusa/backend/internal/server"
	"github.com/nusa/backend/internal/service"
	"github.com/nusa/backend/modules/achievement"
	assessmentHandler "github.com/nusa/backend/modules/assessment"
	authHandler "github.com/nusa/backend/modules/auth"
	curriculumHandler "github.com/nusa/backend/modules/curriculum"
	learningPlanningHandler "github.com/nusa/backend/modules/learning_planning"
	reportingHandler "github.com/nusa/backend/modules/reporting"
	roleHandler "github.com/nusa/backend/modules/roles"
	schoolHandler "github.com/nusa/backend/modules/schools"
	userHandler "github.com/nusa/backend/modules/users"
	jwtService "github.com/nusa/backend/pkg/jwt"
	"go.uber.org/zap"
)

type App struct {
	Config *config.Config
	Logger *logger.Logger
	DB     *db.Postgres
	SQLxDB *sqlx.DB
	Server *server.Server
}

func New() (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	log, err := logger.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	log.Info("Initializing application")

	// Initialize pgxpool for advanced features
	pg, err := db.NewPostgres(&cfg.Database)
	if err != nil {
		log.Error("Failed to connect to database", zap.Error(err))
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Info("Database connected successfully")

	// Initialize sqlx.DB for repositories (compatible with existing repository implementations)
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host,
		cfg.Database.Port, cfg.Database.DBName, cfg.Database.SSLMode,
	)
	sqlxDB, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Error("Failed to connect to database with sqlx", zap.Error(err))
		return nil, fmt.Errorf("failed to connect to database with sqlx: %w", err)
	}
	sqlxDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlxDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlxDB.SetConnMaxLifetime(cfg.Database.ConnMaxLifetime)
	sqlxDB.SetConnMaxIdleTime(cfg.Database.ConnMaxIdleTime)

	// Initialize repositories
	log.Info("Initializing repositories")
	userRepo := repository.NewUserRepository(sqlxDB)
	roleRepo := repository.NewRoleRepository(sqlxDB)
	schoolRepo := repository.NewSchoolRepository(sqlxDB)
	refreshTokenRepo := repository.NewRefreshTokenRepository(sqlxDB)

	// Education domain repositories
	curriculumRepo := repository.NewCurriculumRepository(sqlxDB)
	tpRepo := repository.NewTPRepository(sqlxDB)
	learningPlanningRepo := repository.NewLearningPlanningRepository(sqlxDB)
	assessmentRepo := repository.NewAssessmentRepository(sqlxDB)
	reportingRepo := repository.NewReportingRepository(sqlxDB)

	// Initialize services
	log.Info("Initializing services")
	userService := service.NewUserService(userRepo, roleRepo)
	schoolService := service.NewSchoolService(schoolRepo)
	roleService := service.NewRoleService(roleRepo)

	// Education domain services
	curriculumService := service.NewCurriculumService(curriculumRepo)
	tpService := service.NewTPService(tpRepo)
	atpService := service.NewLearningPlanningService(learningPlanningRepo, learningPlanningRepo)
	modulAjarService := service.NewLearningPlanningService(learningPlanningRepo, learningPlanningRepo)
	assessmentService := service.NewAssessmentService(assessmentRepo)
	achievementService := service.NewAchievementService(assessmentRepo, tpRepo)
	reportingService := service.NewReportingService(reportingRepo, achievementService)

	// Initialize JWT service
	log.Info("Initializing JWT service")
	jwtSvc := jwtService.NewService(
		cfg.JWT.Secret,
		cfg.JWT.Expiration,
		7*24*time.Hour, // 7 days for refresh token
		"nusa-backend",
	)

	// Initialize handlers
	log.Info("Initializing handlers")
	authH := authHandler.NewHandler(userService, refreshTokenRepo, jwtSvc, roleRepo, schoolRepo)
	userH := userHandler.NewHandler(userService, roleRepo, schoolRepo)
	schoolH := schoolHandler.NewHandler(schoolService)
	roleH := roleHandler.NewHandler(roleService)

	// Education domain handlers
	curriculumH := curriculumHandler.NewHandler(curriculumService)
	learningPlanningH := learningPlanningHandler.NewHandler(tpService, atpService, modulAjarService)
	assessmentH := assessmentHandler.NewHandler(assessmentService)
	achievementH := achievement.NewHandler(achievementService)
	reportingH := reportingHandler.NewHandler(reportingService)

	// Initialize router with all handlers
	log.Info("Initializing router with routes")
	r := router.NewRouter(authH, userH, schoolH, roleH, curriculumH, learningPlanningH, assessmentH, achievementH, reportingH, jwtSvc, userRepo, schoolRepo)

	// Create server with configured router
	srv := server.NewWithRouter(cfg, log, r.GetEngine())

	// Setup middleware
	setupMiddleware(srv, log)

	log.Info("Application initialization complete")

	return &App{
		Config: cfg,
		Logger: log,
		DB:     pg,
		SQLxDB: sqlxDB,
		Server: srv,
	}, nil
}

func setupMiddleware(srv *server.Server, log *logger.Logger) {
	router := srv.GetRouter()
	router.Use(middleware.RequestID())
	router.Use(middleware.Recovery())
	router.Use(middleware.Logging(log))
	router.Use(middleware.CORS())
}

func (a *App) Run() error {
	a.Logger.Info("Starting application")

	go func() {
		if err := a.Server.Start(); err != nil {
			a.Logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	a.Logger.Info("Shutting down application...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {
		a.Logger.Error("Server shutdown error", zap.Error(err))
		return err
	}

	if err := a.SQLxDB.Close(); err != nil {
		a.Logger.Error("SQLxDB close error", zap.Error(err))
		return err
	}

	if err := a.DB.Close(); err != nil {
		a.Logger.Error("Database close error", zap.Error(err))
		return err
	}

	if err := a.Logger.Sync(); err != nil {
		return err
	}

	a.Logger.Info("Application shutdown complete")
	return nil
}
