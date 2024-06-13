package main

import (
	"flashcard/internal/card"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type handler struct {
	db *gorm.DB
}

func main() {
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

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/api/card", h.handleCardCreate)
	e.GET("/api/card", h.handleCardIndex)
	e.DELETE("/api/card/:id", h.handleCardDelete)
	e.PUT("/api/card", h.handleCardUpdate)
	e.POST("/api/result", h.handleResultCreate)
	e.GET("/api/play", h.handlePlay)
	e.Logger.Fatal(e.Start(":1323"))
}
