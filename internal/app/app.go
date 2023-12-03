package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/Dorrrke/auth-service/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TODO: инит хранилища (storage)

	//TODO: init auth server (auth)

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
