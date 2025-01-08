package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type PostsStorage struct {
	db *sql.DB
}

type Post struct {
	ID int64 `json:"id"`
	Title string `json:"title" binding:"required"`
	Content string	`json:"content" binding:"required"`
	Tags []string `json:"tags"`
	UserId int64 `json:"user_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}

func (s *PostsStorage) Create(context context.Context, post *Post) error {
	query := `
				INSERT into posts 
				(title, content, tags, user_id) 
				values (?, ?, ?, ?) RETURNING 
				id, created_at, updated_at
			`

	row := s.db.QueryRowContext(context, query, post.Title, post.Content, pq.Array(post.Tags), post.UserId)

	return row.Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
}