package pgrepo

import (
	"context"
	"fmt"
	"url-short/internal/repo"
	"url-short/pkg/pg"

	"github.com/jackc/pgx/v5/pgconn"
)

type urlRepo struct {
	db *pg.PgDb
}

func New(db *pg.PgDb) *urlRepo {
	return &urlRepo{
		db: db,
	}
}


func (u urlRepo) Create(ctx context.Context, url string, alias string) error{
	_, err := u.db.Exec(ctx, "INSERT INTO urls(url, alias) VALUES($1,$2)", url, alias)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == repo.PgCodeDublicate {
				return fmt.Errorf("%s: %w", pgErr.ColumnName, repo.ErrAliasExists)
			}
		}
		fmt.Println("err", err)
		return fmt.Errorf("failed insert url: %w", err)
	}
	return nil
}


func(u urlRepo) GetByAlias(ctx context.Context, alias string) (string, error) {

	var resUrl string

	err := u.db.QueryRow(ctx, "SELECT url FROM urls WHERE alias = $1", alias).Scan(&resUrl)
	if err != nil {
		return "", fmt.Errorf("failed get url: %w", err)
	}
	
	return resUrl, nil
}