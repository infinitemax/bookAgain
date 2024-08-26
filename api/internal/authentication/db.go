package authentication

import (
	"context"
	"fmt"
	"github.com/infinitemax/bookAgain/internal/users"
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

func (db *Db) CreateUser(ctx context.Context, user *users.User) error {

	sql := `insert into books.users (user_name, email, created_at, hashed_password)
			values ($1, $2, $3, $4)`

	_, err := db.pool.Exec(ctx, sql, user.UserName, user.Email, user.CreatedAt, user.HashedPassword)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}
