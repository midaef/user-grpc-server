package service

import (
	"context"
	"github.com/NameLessCorporation/user-grpc-server/internal/models"
	"github.com/NameLessCorporation/user-grpc-server/internal/repository"
	"github.com/NameLessCorporation/user-grpc-server/pkg/helpers"
)

// Users ...
type Users interface {
	SignUp(ctx context.Context, user *models.User) error
}

// Services ...
type Services struct {
	Users Users
}

type Dependencies struct {
	Repository *repository.Repositories
	Hasher     *helpers.Md5
	JWTManager helpers.JWTManager
}

// NewServices ...
func NewServices(deps *Dependencies) *Services {
	usersService := NewUsersService(deps.Repository.Users, deps.Hasher, deps.JWTManager)
	return &Services{
		Users: usersService,
	}
}
