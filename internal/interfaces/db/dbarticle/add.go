package dbarticle

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"newsWebApp/internal/core/article"
)

const insertQuery = `INSERT INTO articles (title, content) VALUES (@title, @content) RETURNING id`
const existCategoryQuery = `INSERT INTO article_categories (article_id, category_id)
							SELECT @article_id, id 
							FROM categories 
							WHERE name=@name`

func (a *DBArticle) Add(ctx context.Context, model article.DBModel) error {
	tx, err := a.db.Begin(ctx)

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	args := pgx.NamedArgs{"title": model.Title, "content": model.Content}
	err = tx.QueryRow(ctx, insertQuery, args).Scan(&model.ID)
	if err != nil {
		return fmt.Errorf("DbArticle.Add: failed to insert dbarticle %w", err)
	}

	for _, v := range model.Categories {
		args = pgx.NamedArgs{"article_id": model.ID, "name": v.Name}
		_, err = tx.Exec(ctx, existCategoryQuery, args)
		if err != nil {
			return fmt.Errorf("DbArticle.Add: operations with category was failed %w", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("DbArticle.Add: failed to commit transaction %w", err)
	}

	return nil
}
