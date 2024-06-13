package main

import (
	"flashcard/internal/card"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) handleCardIndex(c echo.Context) error {
	cards, err := card.ListCards(h.db, c.QueryParam("category"))

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"cards": cards,
	})
}

func (h *handler) handleCardCreate(c echo.Context) error {
	var params struct {
		Category      string
		QuestionText  string
		QuestionImage string
		AnswerText    string
		AnswerImage   string
	}

	if err := c.Bind(&params); err != nil {
		return err
	}

	newCard := card.Card{
		Category:      params.Category,
		QuestionText:  params.QuestionText,
		QuestionImage: params.QuestionImage,
		AnswerText:    params.AnswerText,
		AnswerImage:   params.AnswerImage,
	}
	if err := card.CreateCard(h.db, &newCard); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, newCard)
}

func (h *handler) handleCardDelete(c echo.Context) error {
	var params struct {
		ID uint `param:"id"`
	}

	if err := c.Bind(&params); err != nil {
		return err
	}
	return card.DeleteCard(h.db, params.ID)
}

func (h *handler) handleCardUpdate(c echo.Context) error {
	var params card.CardUpdateParams

	if err := c.Bind(&params); err != nil {
		return err
	}

	card, err := card.CardUpdate(h.db, params)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, card)
}
