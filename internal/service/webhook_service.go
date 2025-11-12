package service

import (
	"context"
	"io"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/rahulSailesh-shah/ch8n_go/internal/db/repo"
	"github.com/rahulSailesh-shah/ch8n_go/internal/dto"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/inngest"
)

type WebhookService interface {
	ProcessWebhook(ctx context.Context, req *dto.WebhookRequest) (*dto.WorkflowResponse, error)
}

type webhookService struct {
	queries        *repo.Queries
	inngestService *inngest.Inngest
	db             *pgxpool.Pool
}

func NewWebhookService(queries *repo.Queries, db *pgxpool.Pool, inngest *inngest.Inngest) WebhookService {
	return &webhookService{
		queries:        queries,
		db:             db,
		inngestService: inngest,
	}
}

func (s *webhookService) ProcessWebhook(ctx context.Context, req *dto.WebhookRequest) (*dto.WorkflowResponse, error) {
	workflow, err := s.queries.GetWorkflow(ctx, req.WorkflowID)
	if err != nil {
		return nil, err
	}

	nodes, err := s.queries.GetNodesByWorkflowID(ctx, workflow.ID)
	if err != nil {
		return nil, err
	}

	edges, err := s.queries.GetConnectionsByWorkflowID(ctx, workflow.ID)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(req.Data)
	if err != nil {
		return nil, err
	}

	s.inngestService.ExecuteWorkflow(ctx, &inngest.ExecuteWorkflowRequest{
		WorkflowID: req.WorkflowID,
		Nodes:      nodes,
		Edges:      edges,
		Data:       bodyBytes,
	})

	return toWorkflowResponse(&workflow, nodes, edges), nil
}
