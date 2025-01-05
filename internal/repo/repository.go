package repo

import (
	"MyDrive/internal/repo/filesystem"
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	ErrDuplicateEmail    = errors.New("a user with this email already exists")
	ErrDuplicateUsername = errors.New("a user with this username already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Repository struct {
	Users interface {
		GetByID(ctx context.Context, userID int64) (*User, error)
		GetByEmail(ctx context.Context, email string) (*User, error)
		Create(ctx context.Context, tx *sql.Tx, user *User) error
		Delete(ctx context.Context, userID int64) error
	}
	FilesSystem interface {
		CreateFile(path string) (err error)
		DeleteFile(path string) (err error)
		UpdateFile(path string, content []byte, updateAt int64) (err error)
		ReadFile(path string) ([]byte, error)
	}
}

func NewRepo(db *sql.DB) Repository {
	return Repository{
		Users:       &userRepo{db},
		FilesSystem: filesystem.New(),
	}
}

func withTransaction(db *sql.DB, ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
