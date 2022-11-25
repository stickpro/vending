package app

import (
	"context"
	"errors"
	"github.com/stickpro/vending/internal/config"
	"github.com/stickpro/vending/internal/repository"
	"github.com/stickpro/vending/internal/router"
	"github.com/stickpro/vending/internal/server"
	"github.com/stickpro/vending/internal/service"
	"github.com/stickpro/vending/pkg/database/pgsql"
	"github.com/stickpro/vending/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(configDir string) {
	cfg, err := config.Init(configDir)

	if err != nil {
		logger.Error(err)
		return
	}
	db, _ := pgsql.ConnectionDataBase(cfg.DB.Host, cfg.DB.Username, cfg.DB.Password, cfg.DB.DBName, cfg.DB.Port)

	repos := repository.NewRepositories(db)
	// migration
	repository.Migrate(repos)

	services := service.NewServices(service.Deps{
		Repository: repos,
	})

	router := router.NewRouter(services)

	srv := server.NewServer(cfg.HTTP, router.Init())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
}
