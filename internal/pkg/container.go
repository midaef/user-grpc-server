package pkg

import (
	"github.com/NameLessCorporation/user-grpc-server/internal/repository"
	"github.com/NameLessCorporation/user-grpc-server/internal/service"
	"go.uber.org/zap"
)

// Container ...
type Container struct {
	Repositories *repository.Repositories
	Services     *service.Services
	Logger       *zap.Logger
}

// NewContainer ...
func NewContainer(repository *repository.Repositories, services *service.Services, logger *zap.Logger) *Container {
	return &Container{
		Repositories: repository,
		Services:     services,
		Logger:       logger,
	}
}
