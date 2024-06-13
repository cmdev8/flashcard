package main

import (
	"flashcard/internal/card"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) handleResultCreate(c echo.Context) error {
	var params struct {
		CardID  uint
		Success bool
	}

	if err := c.Bind(&params); err != nil {
		return err
	}

	if err := card.CreateResult(h.db, params.CardID, params.Success); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}
