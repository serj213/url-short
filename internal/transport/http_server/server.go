package httpserver

import "context"

type urlService interface {
    Create(ctx context.Context, url string, alias string) (string, error)
    GetUrl(ctx context.Context, alias string) (string, error)
}

type HttpServer struct {
    UrlService urlService
}

func New(urlService urlService) *HttpServer{
    return &HttpServer{
        UrlService: urlService,
    }
}