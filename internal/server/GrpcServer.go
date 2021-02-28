package server

import "github.com/NameLessCorporation/user-grpc-server/internal/service"

// GRPCServer ...
type GRPCServer struct {
	Services *service.Services
}

// NewGRPCServer ...
func NewGRPCServer(services *service.Services) *GRPCServer {
	return &GRPCServer{
		Services: services,
	}
}
