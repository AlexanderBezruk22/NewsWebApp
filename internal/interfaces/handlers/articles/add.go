package articles

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"newsWebApp/internal/core/article"
)

type ArticleService interface {
	Add(ctx context.Context, model article.DBModel) error
	Update(ctx context.Context, model article.DBModel) error
	Fetch(ctx context.Context, categories []string) ([]article.DBModel, error)
}

type ArticleHandlers struct {
	service ArticleService
}

func NewArticleHandlers(service ArticleService) *ArticleHandlers {
	return &ArticleHandlers{service: service}
}

type Request struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category []string `json:"category"`
}

func (h *ArticleHandlers) Add(c *fiber.Ctx) error {
	var request Request
	if err := c.BodyParser(&request); err != nil {
		return fiber.ErrBadRequest
	}

	var categories []article.Category
	for _, v := range request.Category {
		categories = append(categories, article.Category{Name: v})
	}

	err := h.service.Add(
		c.Context(),
		article.DBModel{
			Title:      request.Title,
			Content:    request.Content,
			Categories: categories,
		})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Articles added"})
}
