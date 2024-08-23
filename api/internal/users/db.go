package users

import "github.com/jackc/pgx/v5/pgxpool"

type Db struct {
	pool pgxpool.Pool
}

func NewDb(pool pgxpool.Pool) *Db {
	return &Db{
		pool: pool,
	}
}

func (db *Db) CreateUser() error {

	return nil
}
