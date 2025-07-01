package articles

import (
	"context"
	"newsWebApp/internal/core/article"
)

type DBArticles interface {
	Add(ctx context.Context, model article.DBModel) error
	UpdateArticle(ctx context.Context, model article.DBModel) error
	FetchArticles(ctx context.Context, categories []string) ([]article.DBModel, error)
}

type ArticleService struct {
	db DBArticles
}

func NewArticleService(db DBArticles) *ArticleService {
	return &ArticleService{db: db}
}
