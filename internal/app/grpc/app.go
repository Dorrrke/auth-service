package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	authgrpc "github.com/Dorrrke/auth-service/internal/grpc/auth"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(
	log *slog.Logger,
	port int,
) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.RegisterServerAPI(gRPCServer)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"
	log := a.log.With(slog.String("op", op))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return errors.Wrap(err, op)
	}

	log.Info("Starting GRPC Server ", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"
	a.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
