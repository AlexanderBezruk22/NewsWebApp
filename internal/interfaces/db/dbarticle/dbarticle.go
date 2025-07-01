package dbarticle

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type PgxPool interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Close()
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

type DBArticle struct {
	db PgxPool
}

func New(db PgxPool) *DBArticle {
	return &DBArticle{db: db}
}
