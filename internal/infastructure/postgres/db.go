package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBSource interface {
	Connection() string
}

func DB(ctx context.Context, source DBSource) (*pgxpool.Pool, error) {
	var (
		db  *pgxpool.Pool
		err error
	)

	if db, err = pgxpool.New(ctx, source.Connection()); err != nil {
		return nil, err
	}

	if err = db.Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func DBClose(db *pgxpool.Pool) {
	defer db.Close()
}
