package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"newsWebApp/internal/interfaces/db/dbarticle"
	articlesHandlers "newsWebApp/internal/interfaces/handlers/articles"
	"newsWebApp/internal/services/articles"
)

func RegisterRoutes(app *fiber.App, db *pgxpool.Pool) {

	dbArticle := dbarticle.New(db)
	articleService := articles.NewArticleService(dbArticle)
	articleHandlers := articlesHandlers.NewArticleHandlers(articleService)

	app.Post("/news", articleHandlers.Add)
	app.Put("/news/:id", articleHandlers.Update)
	app.Get("/news", articleHandlers.Fetch)
}
