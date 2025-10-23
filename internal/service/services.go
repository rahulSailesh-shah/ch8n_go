package service

import (
	"github.com/rahulSailesh-shah/ch8n_go/internal/db/repo"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/config"
)

type Service struct {
	Workflow WorkflowService
}

func NewService(queries *repo.Queries, cfg *config.AppConfig) *Service {
	return &Service{
		Workflow: NewWorkflowService(queries),
	}
}
