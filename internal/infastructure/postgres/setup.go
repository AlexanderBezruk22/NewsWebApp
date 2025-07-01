package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"newsWebApp/internal/infastructure/postgres/source"
)

func Setup(ctx context.Context, logger *logrus.Logger) (*pgxpool.Pool, error) {
	db, err := DB(ctx, source.WithENV())

	if err != nil {
		return nil, err
	}

	logger.Info("Success to connect postgres with pgx")

	return db, nil
}
