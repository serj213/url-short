package service

import (
	"context"
	"fmt"
	"log/slog"
	"url-short/internal/lib/logger/sl"
	randomalias "url-short/internal/lib/randomAlias"
)

type UrlRepository interface {
	Create(ctx context.Context, url string, alias string) error
	GetByAlias(ctx context.Context, alias string) (string, error)
}

type urlService struct {
	log *slog.Logger
	repo UrlRepository
}

func New(log *slog.Logger, repo UrlRepository) *urlService {
	return &urlService{
		log: log,
		repo: repo,
	}
}

func (r urlService) Create(ctx context.Context, url string, alias string) (string, error) {
	const lengthAlias = 5
	const id = "urlService.Create"
	log := r.log.With(slog.String("op", id))

	aliasRes := alias

	if alias == "" {
		aliasRes = randomalias.RandomAlias(lengthAlias)
	}

	log.Info(aliasRes)

	err := r.repo.Create(ctx, url, aliasRes)
	
	if err != nil {
		log.Error(fmt.Sprintf("failed create alias %s", err.Error()))
		return "", err
	}

	fmt.Println("test err:", err)
	
	return aliasRes, nil
}

func (r urlService) GetUrl(ctx context.Context, alias string) (string, error) {
	const op = "urlService.GetUrl"
	log := r.log.With(
		slog.String("op", op),
		slog.String("alias", alias),
	)

	url, err := r.repo.GetByAlias(ctx, alias)

	if err != nil {
		log.Error("failed get url by alias: %w", sl.Err(err))
		return "", err
	}

	return url, nil
}


