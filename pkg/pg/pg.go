package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgDb struct {
	*pgxpool.Pool
}


func Deal(dsn string) (*PgDb, error) {

	if dsn == "" {
		return nil, fmt.Errorf("dsn is empty")
	}

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("failed connect database: %w", err)
	}


	return &PgDb{
		Pool: pool,
	}, nil

}