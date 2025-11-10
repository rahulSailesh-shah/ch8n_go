package app

import (
	"context"
	"fmt"

	"github.com/rahulSailesh-shah/ch8n_go/internal/db/repo"
	"github.com/rahulSailesh-shah/ch8n_go/internal/service"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/config"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/database"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/inngest"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/registry"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/template"

	// Initialize all nodes
	_ "github.com/rahulSailesh-shah/ch8n_go/pkg/node/native/http_node"
	_ "github.com/rahulSailesh-shah/ch8n_go/pkg/node/native/manual_trigger"
)

type App struct {
	Config  *config.AppConfig
	DB      database.DB
	Service *service.Service
	Inngest *inngest.Inngest
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

	nodeRegistry := registry.NewNodeRegistry()
	if err := registry.InitializeAllNodes(nodeRegistry); err != nil {
		return nil, err
	}

	templateEngine := template.NewTemplateEngine()

	inngestService, err := inngest.NewInngest(nodeRegistry, templateEngine)
	if err != nil {
		return nil, err
	}

	queries := repo.New(dbInstance)
	service := service.NewService(queries, cfg, dbInstance, inngestService)

	return &App{
		Config:  cfg,
		DB:      db,
		Service: service,
		Inngest: inngestService,
	}, nil
}
