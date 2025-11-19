package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetById(context.Context, int64) (*Post, error)
	}
	Users interface {
		Create(context.Context, *User) error
	}
	Comments interface {
		GetByPostID(context.Context, int64) ([]Comment, error)
	}
}

func NewPostgrestStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostsStore{db},
		Users:    &UsersStore{db},
		Comments: &CommentStore{db},
	}
}
