package repository

import (
	"context"

	"github.com/NameLessCorporation/user-grpc-server/internal/models"
	"github.com/jmoiron/sqlx"
)

// UsersRepository ...
type UsersRepository struct {
	db *sqlx.DB
}

// NewUsersRepository ...
func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

// Create ...
func (r *UsersRepository) Create(ctx context.Context, user *models.User) error {
	return nil
}
