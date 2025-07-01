package articles

import (
	"context"
	"fmt"
	"newsWebApp/internal/core/article"
)

func (s *ArticleService) Add(ctx context.Context, model article.DBModel) error {
	err := s.db.Add(ctx, model)
	if err != nil {
		return fmt.Errorf("ArticleService.Add: failed to add dbarticle into db: %w", err)
	}

	return nil
}
