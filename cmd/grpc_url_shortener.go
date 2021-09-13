package cmd

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"grpc-url-shortener/base63"
	grpc "grpc-url-shortener/pkg/protocol/grcp"
	"grpc-url-shortener/pkg/urlshortener/controller"
	"grpc-url-shortener/pkg/urlshortener/repository/postgres_sql"
	"grpc-url-shortener/pkg/urlshortener/service"
	"log"
	"os"
)

type App struct {
	db *sql.DB
}

func RunServer() error {
	ctx := context.Background()
	db, err := newDB()
	if err != nil {
		log.Println("ERROR creating new database: ", err)
		return err
	}

	defer db.Close()

	base63Handler := base63.NewBase63Handler()

	// CREATE REPOSITORY
	repo := postgres_sql.New(db)

	// RUN MIGRATIONS
	if err = repo.Migrate(); err != nil {
		log.Println("ERROR MIGRATING: ", err)
		return err
	}
	//CREATE SERVICE
	urlService := service.New(repo, base63Handler)

	//CREATE CONTROLLER
	urlController := controller.New(urlService)

	return grpc.RunServer(ctx, urlController, os.Getenv("APP_PORT"))
}

func newDB() (*sql.DB, error) {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
