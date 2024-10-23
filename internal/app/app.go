package app

import (
	grpcApp "grpcTrain/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcApp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePAth string,
	tokerTTL time.Duration,
) *App {

	//

	//

	grpcApp := grpcApp.New(log, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
