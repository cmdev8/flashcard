package main

import (
	"flashcard/internal/card"
	"net/http"
	"slices"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func (h *handler) handleCategoryIndex(c echo.Context) error {
	var cards []card.Card
	if err := h.db.Select("category").Find(&cards).Error; err != nil {
		return err
	}

	allCategories := lo.Map(cards, func(item card.Card, _ int) string {
		return item.Category
	})

	result := lo.Uniq(allCategories)

	slices.Sort(result)

	return c.JSON(http.StatusOK, map[string][]string{
		"Categories": result,
	})
}
