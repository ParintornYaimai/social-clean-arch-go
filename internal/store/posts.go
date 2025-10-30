package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostsStore struct {
	db *sql.DB
}

func (s *PostsStore) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO post (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`
	row := s.db.QueryRowContext(ctx, query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
	)

	if err := row.Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt); err != nil {
		return err
	}

	return nil
}

type GetPost struct {
	ID int64 `json:"id"`
}

var (
	ErrNotFound = errors.New("record not found")
)

func (s *PostsStore) GetById(ctx context.Context, id *GetPost) (*Post, error) {
	query := `SELECT * FROM post WHERE id = $id`

	var post Post
	err := s.db.QueryRowContext(ctx, query).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, pq.Array(&post.Tags), &post.CreatedAt, &post.UpdatedAt)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
		}
		return nil, err
	}

	return &post, nil
}
