package grpcApp

import (
	"fmt"
	"google.golang.org/grpc"
	authgRPC "grpcTrain/internal/grpc/auth"
	"log/slog"
	"net"
)

type App struct {
	log         *slog.Logger
	gGRPCServer *grpc.Server
	port        int
}

func New(
	log *slog.Logger,
	port int,
) *App {
	gRPCServer := grpc.NewServer()
	authgRPC.Register(gRPCServer)
	return &App{
		log:         log,
		gGRPCServer: gRPCServer,
		port:        port,
	}
}

func (a *App) Run() error {
	const op = "grpcApp.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", slog.String("addr", l.Addr().String()))

	if err := a.gGRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Stop() {
	const op = "grpcApp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", a.port))

	a.gGRPCServer.GracefulStop()
}
