package main

import (
	"context"
	"embed"
	"errors"
	"flashcard/internal/card"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ardanlabs/conf/v3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const webAppDir = "webapp/dist"

//go:embed webapp/dist
var webappEmbedFs embed.FS

type config struct {
	DBPath     string `conf:"default:test.db"`
	ServerPort int    `conf:"default:8013"`
}

func main() {
	if err := run(); err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}
}

type handler struct {
	db *gorm.DB
}

func run() error {
	var cfg config

	if help, err := conf.Parse("", &cfg); err != nil {
		fmt.Println(help)
		if errors.Is(err, conf.ErrHelpWanted) {
			return nil
		}

		return err
	}

	db, err := gorm.Open(sqlite.Open(cfg.DBPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&card.Card{})
	db.AutoMigrate(&card.Result{})

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

	h := handler{
		db: db,
	}
	registerRoutes(e, &h)

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "/",
		HTML5: true,
		Filesystem: func(path string, embededFiles embed.FS) http.FileSystem {
			fsys, err := fs.Sub(embededFiles, path)
			if err != nil {
				panic(err)
			}

			return http.FS(fsys)
		}(webAppDir, webappEmbedFs),
	}))

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

	return nil
}
