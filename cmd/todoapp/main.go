package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/NeverEverLive/todo-go/internal/core/logger"
	core_pgx_pool "github.com/NeverEverLive/todo-go/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/NeverEverLive/todo-go/internal/core/transport/http/middleware"
	core_http_server "github.com/NeverEverLive/todo-go/internal/core/transport/http/server"
	users_postgres_repository "github.com/NeverEverLive/todo-go/internal/features/users/repository/postgres"
	users_service "github.com/NeverEverLive/todo-go/internal/features/users/service"
	users_transport_http "github.com/NeverEverLive/todo-go/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("Failed to init application logger, error:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("Initializing postgres connection pool")
	pool, err := core_pgx_pool.NewPool(
		ctx,
		core_pgx_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("Failed to init database connection pool, error:", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("Initializing feature!", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService)

	logger.Debug("Initializing HTTP server")

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouterV1 := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersionV1)
	apiVersionRouterV1.RegisterRouter(usersTransportHTTP.Routes()...)
	httpServer.RegisterAPIRouters(apiVersionRouterV1)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
