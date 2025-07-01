package dbarticle

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"newsWebApp/internal/core/article"
)

//const fetchCategoriesQuery = `SELECT c.name
//							  FROM categories c
//							  INNER JOIN article_categories ac
//							     ON c.id = ac.category_id
//							  WHERE ac.article_id=@article_id;`

const fetchWithFilterQuery = `
							SELECT DISTINCT a.id, a.title, a.content 
							FROM articles a 
							INNER JOIN article_categories ac ON a.id = ac.article_id 
							INNER JOIN categories c ON ac.category_id = c.id 
							WHERE c.name = ANY(@categories);`

const fetchWithoutFilterQuery = `SELECT * FROM articles;`

func (a *DBArticle) FetchArticles(ctx context.Context, categories []string) ([]article.DBModel, error) {
	tx, err := a.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("DBArticle.Fetch: error starting transaction: %v", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	var results []article.DBModel
	if len(categories) == 0 {
		results, err = a.FetchWithoutFilter(ctx, tx)
	} else {
		results, err = a.FetchWithFilter(ctx, tx, categories)
	}
	if err != nil {
		return nil, fmt.Errorf("DBArticle.Fetch: error fetching articles: %v", err)
	}

	return results, nil
}

func (a *DBArticle) FetchWithFilter(ctx context.Context, tx pgx.Tx, categories []string) ([]article.DBModel, error) {
	rows, err := tx.Query(ctx, fetchWithFilterQuery, pgx.NamedArgs{"categories": categories})
	if err != nil {
		return nil, fmt.Errorf("DBArticle.FetchWithoutFilter: failed to fetch articles: %w", err)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("DBArticle.FetchCategories: rows return error: %w", err)
	}
	defer func() {
		rows.Close()
	}()

	articles, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[article.DBModel])
	if err != nil {
		return nil, fmt.Errorf("DBArticle.FetchWithoutFilter: failed to collect rows: %w", err)
	}

	return articles, nil
}

func (a *DBArticle) FetchWithoutFilter(ctx context.Context, tx pgx.Tx) ([]article.DBModel, error) {
	rows, err := tx.Query(ctx, fetchWithoutFilterQuery)
	if err != nil {
		return nil, fmt.Errorf("DBArticle.FetchWithoutFilter: failed to fetch articles: %w", err)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("DBArticle.FetchCategories: rows return error: %w", err)
	}
	defer func() {
		rows.Close()
	}()

	articles, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[article.DBModel])
	if err != nil {
		return nil, fmt.Errorf("DBArticle.FetchWithoutFilter: failed to collect rows: %w", err)
	}

	return articles, nil
}

// Если понадобится получение категорий
//func (a *DBArticle) FetchCategories(ctx context.Context, tx pgx.Tx,articleID int) ([]article.Category, error) {
//	rows, err := tx.Query(ctx, fetchCategoriesQuery, pgx.NamedArgs{"article_id": articleID})
//	if err != nil {
//		return nil, fmt.Errorf("DBArticle.FetchCategories: failed to fetch categories of article: %w", err)
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, fmt.Errorf("DBArticle.FetchCategories: rows return error: %w", err)
//	}
//	defer func() {
//		rows.Close()
//	}()
//
//	var categories []article.Category
//	categories, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[article.Category])
//	if err != nil {
//		return nil, fmt.Errorf("DBArticle.FetchCategories: failed to collect rows of categories %w", err)
//	}
//
//	return categories, nil
//}
