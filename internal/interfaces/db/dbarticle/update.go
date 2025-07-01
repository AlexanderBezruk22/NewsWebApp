package dbarticle

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"newsWebApp/internal/core/article"
	"newsWebApp/internal/core/domainerror"
	"sort"
	"strings"
)

const deleteCategoryQuery = `DELETE FROM article_categories 
       		  				 WHERE article_id=@article_id;`

func (a *DBArticle) UpdateArticle(ctx context.Context, model article.DBModel) error {
	fields := checkFields(model)
	if len(fields) == 0 {
		return nil
	}

	keys := make([]string, 0, len(fields))
	for key := range fields {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	setClauses := make([]string, 0, len(keys))

	for _, key := range keys {
		setClauses = append(setClauses, fmt.Sprintf("%s=@%s", key, key))
	}

	fields["id"] = model.ID

	tx, err := a.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("DBArticle.UpdateArticle begin failed: %v", err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	var lockedID string
	lockQuery := "SELECT id FROM articles WHERE id = @id FOR UPDATE NOWAIT"
	err = tx.QueryRow(ctx, lockQuery, pgx.NamedArgs{"id": model.ID}).Scan(&lockedID)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "55P03" { // Код ошибки "lock_not_available"
		return domainerror.ErrTooManyRequests
	}

	if err != nil {
		return fmt.Errorf("DBArticle.UpdateArticle lock failed: %w", err)
	}

	query := fmt.Sprintf("UPDATE articles SET %s WHERE id=@id", strings.Join(setClauses, ", "))
	_, err = tx.Exec(ctx, query, pgx.NamedArgs(fields))
	if err != nil {
		return fmt.Errorf("DBArticle.UpdateArticle: update articles failed: %w", err)
	}

	if len(model.Categories) != 0 {
		err = a.UpdateCategories(ctx, tx, model)
	}

	if err != nil {
		return fmt.Errorf("DBArticle.UpdateArticle failed: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("DBArticle.UpdateArticle commit failed: %v", err)
	}

	return nil
}

func (a *DBArticle) UpdateCategories(ctx context.Context, tx pgx.Tx, model article.DBModel) error {
	args := pgx.NamedArgs{"article_id": model.ID}
	_, err := tx.Exec(ctx, deleteCategoryQuery, args)
	if err != nil {
		return fmt.Errorf("DbArticle.UpdateCategory: delete category from dbarticle was failed %w", err)
	}

	for _, v := range model.Categories {
		args = pgx.NamedArgs{"article_id": model.ID, "name": v.Name}
		_, err = tx.Exec(ctx, existCategoryQuery, args)
		if err != nil {
			return fmt.Errorf("DbArticle.UpdateCategory: operations with category was failed %w", err)
		}

	}

	return nil
}

func checkFields(updateInput article.DBModel) map[string]interface{} {
	fields := make(map[string]interface{})

	if updateInput.Title != "" {
		fields["title"] = updateInput.Title
	}

	if updateInput.Content != "" {
		fields["content"] = updateInput.Content
	}

	return fields
}
