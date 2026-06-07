package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/config"
	"github.com/nusa/backend/internal/logger"
	"go.uber.org/zap"
)

type Server struct {
	router *gin.Engine
	config *config.Config
	logger *logger.Logger
	server *http.Server
}

// New creates a new server with a default router (for backward compatibility)
func New(cfg *config.Config, log *logger.Logger) *Server {
	if cfg.Server.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	return &Server{
		router: router,
		config: cfg,
		logger: log,
	}
}

// NewWithRouter creates a new server with a pre-configured router
func NewWithRouter(cfg *config.Config, log *logger.Logger, router *gin.Engine) *Server {
	return &Server{
		router: router,
		config: cfg,
		logger: log,
	}
}

// SetupRoutes sets up basic health/ready/version routes (for backward compatibility)
// NOTE: When using NewWithRouter, routes are already configured in the router
func (s *Server) SetupRoutes() {
	s.router.GET("/health", s.healthHandler)
	s.router.GET("/ready", s.readyHandler)
	s.router.GET("/version", s.versionHandler)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

func (s *Server) readyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ready",
	})
}

func (s *Server) versionHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": "1.0.0",
		"name":    "NUSA Backend API",
	})
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:         s.config.Server.Port,
		Handler:      s.router,
		ReadTimeout:  s.config.Server.ReadTimeout,
		WriteTimeout: s.config.Server.WriteTimeout,
		IdleTimeout:  60 * time.Second,
	}

	s.logger.Info("Starting server",
		zap.String("port", s.config.Server.Port),
		zap.String("environment", s.config.Server.Environment),
	)

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down server...")

	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}

	s.logger.Info("Server stopped gracefully")
	return nil
}

func (s *Server) GetRouter() *gin.Engine {
	return s.router
}
