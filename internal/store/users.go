package store

import (
	"context"
	"database/sql"
	"time"
)

type UsersStorage struct {
	db *sql.DB
}

type User struct {
	ID int64 `json:"id"`
	Email string `json:"email" binding:"required"`
	UserName string `json:"username"`
	Password string `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *UsersStorage) Create(context context.Context, user *User) error {
	query := `
		INSERT INTO users (
		email, username, password,
		) values (?, ?, ?) RETURNING
		id, created_at, updated_at
	`

	row := s.db.QueryRowContext(context, query, user.Email, user.UserName, user.Password)

	err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	return err
}