package postgres_sql

import (
	"context"
	"database/sql"
	"errors"
	"grpc-url-shortener/pkg/urlshortener/model"
	"grpc-url-shortener/pkg/urlshortener/repository"
	"log"
)

type Postgress struct {
	db *sql.DB
}

func New(db *sql.DB) repository.Repository {
	return Postgress{
		db: db,
	}
}

func (p Postgress)Create(ctx context.Context, shortner model.UrlShortner) (shortURL string, err error) {
	sqlStatement := `INSERT INTO short_links (original_url, short_url, key_for_decode) VALUES ($1, $2, $3) RETURNING short_url`
	if err = p.db.QueryRow(sqlStatement, shortner.OriginalURL, shortner.ShortURL, shortner.Key).Scan(&shortURL);err != nil {
		return
	}
	return
}

func (p Postgress)Get(ctx context.Context, shortURL string) (originalURL string , err error){
	log.Println("short get: ", shortURL)
	sqlStatement := `SELECT original_url FROM short_links WHERE short_url = $1;`
	row := p.db.QueryRow(sqlStatement, shortURL)
	switch err = row.Scan(&originalURL); err {
	case sql.ErrNoRows:
		return "", errors.New("record not found")
	case nil:
		return
	default:
		return
	}
}

func (p Postgress)Migrate() error {
	if _, err := p.db.Exec(`CREATE TABLE IF NOT EXISTS short_links (
  			id SERIAL PRIMARY KEY,
  			original_url TEXT, 
  			short_url TEXT, 
  			key_for_decode TEXT
		);
	`); err != nil{
		return err
	}
	return nil
}