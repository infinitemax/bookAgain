package users

import (
	"context"
	"log"
)

func (db *Db) GetUsers(ctx context.Context) ([]*User, error) {

	sql := `select * from books.users where deleted_at is null`

	rows, err := db.pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersSlice []*User

	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.UserName, &u.Email, &u.CreatedAt, &u.DeletedAt)
		if err != nil {
			log.Fatal(err)
		}
		usersSlice = append(usersSlice, &u)
	}

	return usersSlice, nil
}
