package articles

import (
	"github.com/gofiber/fiber/v2"
)

func (h *ArticleHandlers) Fetch(c *fiber.Ctx) error {
	categories := c.Context().QueryArgs().PeekMulti("categories")
	var categoryList []string
	for _, category := range categories {
		categoryList = append(categoryList, string(category))
	}

	models, err := h.service.Fetch(c.Context(), categoryList)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"articles": models})
}
