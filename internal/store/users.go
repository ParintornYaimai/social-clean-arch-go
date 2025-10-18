package store

import (
	"context"
	"database/sql"
)

type (
	User struct {
		ID        int64  `json:"id"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"-"`
		CreatedAt string `json:"created_at"`
	}

	UsersStore struct {
		db *sql.DB
	}
)

func (u *UsersStore) Create(ctx context.Context, user *User) error {
	query := `INSERT INTO users (username, password, email,) VALUES($1, $2, $3) RETURNING id, created_at`

	row := u.db.QueryRowContext(ctx, query, &user.Username, &user.Password, &user.Email)

	if err := row.Scan(&user.ID, &user.CreatedAt); err != nil {
		return err
	}

	return nil
}
