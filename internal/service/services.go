package service

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rahulSailesh-shah/ch8n_go/internal/db/repo"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/config"
)

type Service struct {
	Workflow WorkflowService
}

func NewService(queries *repo.Queries, cfg *config.AppConfig, db *pgxpool.Pool) *Service {
	return &Service{
		Workflow: NewWorkflowService(queries, db),
	}
}
