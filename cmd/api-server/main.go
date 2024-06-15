package main

import (
	"context"
	"flashcard/internal/card"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type handler struct {
	db *gorm.DB
}

type config struct {
	DBPath     string
	ServerPort int
}

func main() {
	cfg := config{
		DBPath:     "test.db",
		ServerPort: 8013,
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&card.Card{})
	db.AutoMigrate(&card.Result{})

	h := handler{
		db: db,
	}

	e := echo.New()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))

	e.Use(middleware.Recover())

	e.POST("/api/card", h.handleCardCreate)
	e.GET("/api/card", h.handleCardIndex)
	e.DELETE("/api/card/:id", h.handleCardDelete)
	e.PUT("/api/card", h.handleCardUpdate)
	e.POST("/api/result", h.handleResultCreate)
	e.GET("/api/practice", h.handlePractice)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(fmt.Sprintf(":%d", cfg.ServerPort)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
