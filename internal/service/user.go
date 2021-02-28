package service

import (
	"context"
	"github.com/NameLessCorporation/user-grpc-server/internal/models"
	"time"

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
func (s *UsersService) SignUp(ctx context.Context, input *UserSignUpInput) error {
	pass, err := s.hasher.NewMD5Hash(input.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		User: input.User,
		Name: input.Name,
		Password: pass,
		RegisteredAt: time.Now(),
	}
	if err := s.repository.Create(ctx, user); err != nil {
		return err
	}
	return nil
}
