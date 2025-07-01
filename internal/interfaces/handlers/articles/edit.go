package articles

import (
	"github.com/gofiber/fiber/v2"
	"newsWebApp/internal/core/article"
	"strconv"
)

type updateRequest struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category []string `json:"category"`
}

func (h *ArticleHandlers) Update(c *fiber.Ctx) error {
	var req updateRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	articleId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid id"})
	}

	var categories []article.Category
	for _, v := range req.Category {
		categories = append(categories, article.Category{Name: v})
	}

	err = h.service.Update(c.Context(), article.DBModel{
		ID:         articleId,
		Title:      req.Title,
		Content:    req.Content,
		Categories: categories,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Articles updated"})
}
