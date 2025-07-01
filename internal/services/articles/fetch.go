package articles

import (
	"context"
	"fmt"
	"newsWebApp/internal/core/article"
)

func (s *ArticleService) Fetch(ctx context.Context, categories []string) ([]article.DBModel, error) {
	models, err := s.db.FetchArticles(ctx, categories)
	if err != nil {
		return nil, fmt.Errorf("ArticleService.Fetch: failed fetch articles: %w", err)
	}

	return models, nil
}
