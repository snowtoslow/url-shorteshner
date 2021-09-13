package service

import (
	"context"
	"errors"
	"grpc-url-shortener/base63"
	"grpc-url-shortener/pkg/urlshortener/model"
	"grpc-url-shortener/pkg/urlshortener/repository"
	"grpc-url-shortener/utils"
	"net/url"
)

type Service struct {
	repo repository.Repository
	encoder base63.Base63
}

func New(repo repository.Repository, encoder base63.Base63) Service {
	return Service{
		repo: repo,
		encoder: encoder,
	}
}

func (s Service) Create(ctx context.Context, originalURL string) (string, error) {
	uri, err := url.ParseRequestURI(originalURL)
	if err != nil {
		return "", errors.New("invalid url")
	}

	randomNr, stringValue, err := utils.RandomUINT64()
	if err != nil{
		return "", err
	}

	encoded, err := s.encoder.Encode(randomNr)
	if err != nil {
		return "", err
	}

	savedShortURL, err := s.repo.Create(ctx, model.UrlShortner{
		OriginalURL: originalURL,
		ShortURL:    encoded,
		Key:         stringValue,
	})
	if err != nil {
		return "", err
	}

	u := url.URL{Scheme: uri.Scheme, Host:   uri.Host, Path:   savedShortURL}

	return u.String(), nil
}

func (s Service) Get(ctx context.Context, shortURL string) (string, error) {
	uri, err := url.ParseRequestURI(shortURL)
	if err != nil {
		return "", errors.New("invalid url")
	}
	originalURL, err := s.repo.Get(ctx, uri.Path[1:])
	if err != nil{
		return "", err
	}
	return originalURL, nil
}