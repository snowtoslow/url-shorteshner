package postgres_sql

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"grpc-url-shortener/pkg/urlshortener/model"
	"log"
	"testing"
)

var urlShortnerModel = model.UrlShortner{
	OriginalURL: "originalURL",
	Key: "bigUINT64.ToString",
	ShortURL: "shortOfOriginalURL",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCreate(t *testing.T) {
	t.Parallel()
	db, mock := NewMock()
	ctx := context.Background()
	r := require.New(t)
	repo := New(db)

	defer db.Close()

	query := "INSERT INTO short_links (original_url, short_url, key_for_decode) VALUES ($?, $?, %?) RETURNING short_url"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(urlShortnerModel.OriginalURL, urlShortnerModel.ShortURL, urlShortnerModel.Key).WillReturnResult(sqlmock.NewResult(0, 1))


	_, err := repo.Create(ctx, urlShortnerModel)
	r.Error(err)
}

func TestGet(t *testing.T)  {
	t.Parallel()
	ctx := context.Background()
	r := require.New(t)
	db, mock := NewMock()
	repo := New(db)

	defer db.Close()


	query := "SELECT original_url FROM short_links WHERE short_url = $?"

	rows := sqlmock.NewRows([]string{"original_url"}).
		AddRow(urlShortnerModel.OriginalURL)


	mock.ExpectQuery(query).WithArgs(urlShortnerModel.ShortURL).WillReturnRows(rows)

	originalURL, err := repo.Get(ctx, urlShortnerModel.ShortURL)
	r.NoError(err)
	r.NotEmpty(originalURL)
	r.Equal(originalURL, urlShortnerModel.OriginalURL)
}