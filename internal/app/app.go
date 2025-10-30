package app

import (
	"context"
	"fmt"

	"github.com/rahulSailesh-shah/ch8n_go/internal/db/repo"
	"github.com/rahulSailesh-shah/ch8n_go/internal/service"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/config"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/database"
)

type App struct {
	Config  *config.AppConfig
	DB      database.DB
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

	queries := repo.New(dbInstance)
	service := service.NewService(queries, cfg, dbInstance)

	return &App{
		Config:  cfg,
		DB:      db,
		Service: service,
	}, nil
}
