package main

import (
	"flashcard/internal/card"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *handler) handlePlay(c echo.Context) error {
	var params struct {
		Category string
	}

	if err := c.Bind(&params); err != nil {
		return err
	}

	nextCard, err := card.GetNextCard(h.db, params.Category, time.Now())
	if err != nil {
		return err
	}

	if nextCard == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status": "not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Card":   nextCard,
		"Status": "ok",
	})
}
