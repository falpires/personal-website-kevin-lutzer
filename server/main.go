package main

import (
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinlutzer/personal-website/server/pkg/blog"
	"github.com/kevinlutzer/personal-website/server/pkg/healthcheck"
	"github.com/kevinlutzer/personal-website/server/pkg/server"
	"github.com/kevinlutzer/personal-website/server/pkg/visitor"
	cron "github.com/robfig/cron/v3"
	"go.uber.org/zap"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	StaticDir = "../public"
)

func setupDB(logger *zap.Logger) *gorm.DB {

	logger.Sugar().Info("Setting up database connection...")

	if (DBHost == "") || (DBName == "") || (DBPassword == "") || (DBUser == "") {
		os.Exit(ErrMissingDBEnvVars)
	}

	dsn := "host=" + DBHost + " user=" + DBUser + " password=" + DBPassword + " dbname=" + DBName + " port=" + DBPort + " sslmode=" + DBSSLMode + " TimeZone=UTC"
	gorm, err := gorm.Open(gormpostgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		os.Exit(ErrFailedToSetupDB)
	}

	return gorm
}

func setupPing(log *zap.Logger) {
	if PingHost == "" {
		log.Info("No PING_HOST specific so cron job was not setup...")
		return
	}

	log.Info("Setup the ping cron job...")
	cron := cron.New()
	cron.AddFunc("@every 10s", func() {
		if _, err := http.Get("https://" + PingHost); err != nil {
			log.Error("Failed to ping healthcheck", zap.Error(err))
			return
		}

		log.Info("Successfully pinged healthcheck")
	})

	// Starts it's own go routine
	cron.Start()
}

func main() {
	var logger *zap.Logger
	var err error

	if Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
		logger, err = zap.NewProduction()
	} else {
		gin.SetMode(gin.DebugMode)
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		os.Exit(ErrFailedToSetupLogger)
	}

	setupPing(logger)
	if Port == "" {
		os.Exit(ErrPortIsNotSpecified)
	}

	db := setupDB(logger)

	visitorRepo := visitor.NewRepo(db)
	visitorService := visitor.NewService((visitorRepo))

	healthCheckService := healthcheck.NewService(db)

	blogRepo := blog.NewRepo(db)
	blogService := blog.NewService(blogRepo)

	logger.Info("Starting server...")

	s := server.NewServer(StaticDir, Version, Burst, RateLimit, logger, healthCheckService, blogService, visitorService)
	if err := s.Run(":" + Port); err != nil {
		os.Exit(ErrFailedToStartServer)
	}
}
