package articles

import (
	"context"
	"fmt"
	"newsWebApp/internal/core/article"
)

func (s *ArticleService) Update(ctx context.Context, model article.DBModel) error {
	err := s.db.UpdateArticle(ctx, model)
	if err != nil {
		return fmt.Errorf("ArticleService.Update: update dbarticle err: %v", err)
	}

	return nil
}
