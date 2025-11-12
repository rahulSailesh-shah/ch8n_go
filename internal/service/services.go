package service

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rahulSailesh-shah/ch8n_go/internal/db/repo"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/config"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/inngest"
)

type Service struct {
	Workflow WorkflowService
	Webhook  WebhookService
}

func NewService(queries *repo.Queries, cfg *config.AppConfig, db *pgxpool.Pool, inngestService *inngest.Inngest) *Service {
	return &Service{
		Workflow: NewWorkflowService(queries, db, inngestService),
		Webhook:  NewWebhookService(queries, db, inngestService),
	}
}
