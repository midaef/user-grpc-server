package server

import (
	"context"
	"github.com/NameLessCorporation/user-grpc-server/internal/service"

	"github.com/NameLessCorporation/user-grpc-server/internal/api"
)

// SignUp ...
func (s *GRPCServer) SignUp(ctx context.Context, req *api.SignUpRequest) (*api.SignUpResponse, error) {
	user := &service.UserSignUpInput{
		User: req.User,
		Name: req.Name,
		Password: req.Password,
	}

	err := s.Services.Users.SignUp(ctx, user)
	if err != nil {
		return nil, err
	}

	return &api.SignUpResponse{}, nil
}
