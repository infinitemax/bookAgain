package books

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Db struct {
	pool *pgxpool.Pool
}

func NewDb(pool *pgxpool.Pool) *Db {
	return &Db{
		pool: pool,
	}
}

func (db *Db) NewBook(ctx context.Context, book *Book, user uuid.UUID) error {
	sql := `insert into books.books (title, author, user_id)
			values ($1, $2, $3)`

	_, err := db.pool.Exec(ctx, sql, book.Title, book.Author, user)
	if err != nil {
		return err
	}

	return nil
}
