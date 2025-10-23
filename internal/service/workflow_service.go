package service

import (
	"context"

	"github.com/rahulSailesh-shah/ch8n_go/internal/db/repo"
)

type WorkflowService interface {
	CreateWorkflow(ctx context.Context, workflow *repo.Workflow) (*repo.Workflow, error)
	GetWorkflowsByUserID(ctx context.Context, userID string) ([]repo.Workflow, error)
}

type workflowService struct {
	queries *repo.Queries
}

func NewWorkflowService(queries *repo.Queries) WorkflowService {
	return &workflowService{
		queries: queries,
	}
}

func (s *workflowService) CreateWorkflow(ctx context.Context, workflow *repo.Workflow) (*repo.Workflow, error) {
	newWorkflow, err := s.queries.CreateWorkflow(ctx, repo.CreateWorkflowParams{
		Name:        workflow.Name,
		Description: workflow.Description,
		UserID:      workflow.UserID,
	})
	if err != nil {
		return nil, err
	}
	return &newWorkflow, nil
}

func (s *workflowService) GetWorkflowsByUserID(ctx context.Context, userID string) ([]repo.Workflow, error) {
	workflows, err := s.queries.GetWorkflowsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return workflows, nil
}
