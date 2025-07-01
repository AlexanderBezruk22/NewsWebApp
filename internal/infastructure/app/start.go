package app

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
	"net/http"
	"newsWebApp/internal/config"
	"newsWebApp/internal/infastructure/app/server"
	"newsWebApp/internal/infastructure/postgres"
	"newsWebApp/internal/interfaces/handlers"
	"os"
	"os/signal"
	"time"
)

func Start() {
	config.LoadENV()
	debugLogger := NewLogger(logrus.DebugLevel)

	serv := server.New(fiber.Config{
		ServerHeader:  "NewsAPI",
		CaseSensitive: true,
		StrictRouting: true,
	}, ":8080")

	serv.App.Use(logger.New())
	serv.App.Use(compress.New())
	serv.App.Use(limiter.New())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := postgres.Setup(ctx, debugLogger)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	handlers.RegisterRoutes(serv.App, db)

	go func() {
		servErr := serv.Run()
		if errors.Is(servErr, http.ErrServerClosed) {
			debugLogger.Panic(servErr)
		}

		debugLogger.Info("Server successfully stopped")
	}()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt)
	<-termChan

	debugLogger.Info("Gracefully shutting down...")

	err = serv.Shutdown(serv.App)
	if err != nil {
		debugLogger.Panic(err)
	}
	debugLogger.Info("Server successfully stopped")
}

func NewLogger(level logrus.Level) *logrus.Logger {
	debugLogger := logrus.New()
	debugLogger.SetLevel(level)
	return debugLogger
}
