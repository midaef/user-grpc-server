package app

import (
	"github.com/NameLessCorporation/user-grpc-server/internal/models"
	"github.com/NameLessCorporation/user-grpc-server/pkg/helpers"
	"go.uber.org/zap"
	"net"

	"github.com/NameLessCorporation/user-grpc-server/internal/api"
	"github.com/NameLessCorporation/user-grpc-server/internal/pkg"
	"github.com/NameLessCorporation/user-grpc-server/internal/repository"
	"github.com/NameLessCorporation/user-grpc-server/internal/server"
	"github.com/NameLessCorporation/user-grpc-server/internal/service"
	"google.golang.org/grpc"
)

// Run ...
func Run(config *models.Config) {
	logger, err := pkg.ConfigureLogger(config.Logger.LogLevel)
	if err != nil {
		logger.Error("configure logger error \n%v", zap.Error(err))
	}

	logger.Info("Started user-microservice", zap.String("port", config.Server.Port))

	conn := pkg.NewConnection(config.Database.Uri)
	err = conn.Open()
	if err != nil {
		logger.Error("configure connection error", zap.Error(err))
	}

	defer conn.DB.Close()

	logger.Debug("Connected to user-microservice database")

	repositories := repository.NewRepositories(conn.DB)
	hasher := helpers.NewHasher(config.Token.Salt)

	jwtManager, err := helpers.NewPrivateKey(config.Token.SecretKey)
	if err != nil {
		logger.Error("jwt manager error", zap.Error(err))
	}

	deps := &service.Dependencies{
		Repository: repositories,
		Hasher: hasher,
		JWTManager: jwtManager,
	}
	services := service.NewServices(deps)

	container := pkg.NewContainer(repositories, services, logger)

	s := grpc.NewServer()
	srv := server.NewGRPCServer(container.Services)
	api.RegisterSignUpServer(s, srv)

	l, err := net.Listen("tcp", config.Server.Port)
	if err != nil {
		logger.Error("listen tcp error", zap.Error(err))
	}

	err = s.Serve(l)
	if err != nil {
		logger.Error("serve error", zap.Error(err))
	}
}
