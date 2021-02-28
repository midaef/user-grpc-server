package server

import (
	"context"

	"github.com/NameLessCorporation/user-grpc-server/internal/api"
	"github.com/NameLessCorporation/user-grpc-server/internal/models"
)

// SignUp ...
func (s *GRPCServer) SignUp(ctx context.Context, req *api.SignUpRequest) (*api.SignUpResponse, error) {
	user := &models.User{
		Phone: req.Phone,
		Email: req.Email,
		Name:  req.Name,
	}

	err := s.Services.Users.SignUp(ctx, user)
	if err != nil {
		return nil, err
	}

	return &api.SignUpResponse{}, nil
}
