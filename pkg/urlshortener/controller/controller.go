package controller

import (
	"context"
	v1 "grpc-url-shortener/pkg/api/v1"
	"grpc-url-shortener/pkg/urlshortener/service"
)

type Controller struct {
	service service.Service
}

func New(service service.Service) Controller {
	return Controller{
		service: service,
	}
}

func (c Controller)Create(ctx context.Context, request *v1.CreateRequest)(*v1.CreateResponse, error){
	shortURL, err := c.service.Create(ctx,request.Url)
	if err != nil{
		//TODO: wrap error into protobuf error
		return nil, err
	}

	return &v1.CreateResponse{
		ShortUrl: shortURL,
	}, nil
}

func (c Controller)Get(ctx context.Context, request *v1.GetRequest) (*v1.GetResponse, error) {
	originalURL , err := c.service.Get(ctx, request.ShortURL)
	if err != nil{
		//TODO: wrap error into protobuf error
		return nil, err
	}

	return &v1.GetResponse{
		OriginalURL: originalURL,
	}, nil
}
