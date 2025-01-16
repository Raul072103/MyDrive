package repo

import (
	"context"
	"database/sql"
)

type UsersMock struct {
}

func (u *UsersMock) GetByID(ctx context.Context, userID int64) (*User, error) {
	return nil, nil
}

func (u *UsersMock) GetByEmail(ctx context.Context, email string) (*User, error) {
	return nil, nil
}

func (u *UsersMock) Create(ctx context.Context, tx *sql.Tx, user *User) error {
	return nil
}

func (u *UsersMock) Delete(ctx context.Context, userID int64) error {
	return nil
}

func NewUsersMock() *UsersMock {
	return &UsersMock{}
}
