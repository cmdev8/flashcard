package main

import "github.com/labstack/echo/v4"

func registerRoutes(e *echo.Echo, h *handler) {
	e.POST("/api/card", h.handleCardCreate)
	e.GET("/api/card", h.handleCardIndex)
	e.DELETE("/api/card/:id", h.handleCardDelete)
	e.PUT("/api/card", h.handleCardUpdate)
	e.POST("/api/result", h.handleResultCreate)
	e.GET("/api/practice", h.handlePractice)
	e.GET("/api/categories", h.handleCategoryIndex)
}
