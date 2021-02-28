package repository

import (
	"context"

	"github.com/NameLessCorporation/user-grpc-server/internal/models"
	"github.com/jmoiron/sqlx"
)

// Users ...
type Users interface {
	Create(ctx context.Context, user *models.User) error
}

// Repositories ...
type Repositories struct {
	Users Users
}

// NewRepositories ...
func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Users: NewUsersRepository(db),
	}
}
