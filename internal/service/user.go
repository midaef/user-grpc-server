package service

import (
	"context"

	"github.com/NameLessCorporation/user-grpc-server/internal/models"
	"github.com/NameLessCorporation/user-grpc-server/internal/repository"
	"github.com/NameLessCorporation/user-grpc-server/pkg/helpers"
)

// UsersService ...
type UsersService struct {
	repository   repository.Users
	hasher       *helpers.Md5
	tokenManager helpers.JWTManager
}

// NewUsersService ...
func NewUsersService(repository repository.Users, hasher *helpers.Md5, tokenManager helpers.JWTManager) *UsersService {
	return &UsersService{
		repository: repository,
		hasher: hasher,
		tokenManager: tokenManager,
	}
}

// SignUp ...
func (s *UsersService) SignUp(ctx context.Context, user *models.User) error {
	if err := s.repository.Create(ctx, user); err != nil {
		return err
	}
	return nil
}
