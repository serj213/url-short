package httpserver

import "context"

type UrlService interface {
    Create(ctx context.Context, url string, alias string) (string, error)
    GetUrl(ctx context.Context, alias string) (string, error)
}

type HttpServer struct {
    UrlService UrlService
}

func New(urlService UrlService) *HttpServer{
    return &HttpServer{
        UrlService: urlService,
    }
}