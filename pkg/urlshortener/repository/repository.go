package repository

import (
	"context"
	"grpc-url-shortener/pkg/urlshortener/model"
)

type Repository interface {
	Create(ctx context.Context, shortner model.UrlShortner) (shortURL string, err error)
	Get(ctx context.Context, shortURL string) (originalURL string , err error)
	Migrate() error
}
