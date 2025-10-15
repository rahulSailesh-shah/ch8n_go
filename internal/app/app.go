package app

import (
	"context"
	"fmt"

	"github.com/rahulSailesh-shah/ch8n_go/internal/db/sqlc"
	"github.com/rahulSailesh-shah/ch8n_go/internal/repo"
	"github.com/rahulSailesh-shah/ch8n_go/internal/service"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/config"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/database"
)

type App struct {
	Config  *config.AppConfig
	DB      database.DB
	Repo    *repo.Repositories
	Service *service.Service
}

func NewApp(ctx context.Context, cfg *config.AppConfig) (*App, error) {
	db := database.NewDB(ctx, cfg.DB)
	if err := db.Connect(); err != nil {
		return nil, err
	}

	dbInstance := db.GetDB()
	if dbInstance == nil {
		return nil, fmt.Errorf("database instance is nil")
	}

	queries := sqlc.New(dbInstance)
	repo := repo.NewRepositories(queries)
	service := service.NewService(repo, cfg)

	return &App{
		Config:  cfg,
		DB:      db,
		Repo:    repo,
		Service: service,
	}, nil
}
