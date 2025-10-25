package service

import (
	"context"
	"time"

	"github.com/rahulSailesh-shah/ch8n_go/internal/db/repo"
	"github.com/rahulSailesh-shah/ch8n_go/internal/dto"
)

type WorkflowService interface {
	CreateWorkflow(ctx context.Context, workflow *dto.CreateWorkflowRequest) (*dto.WorkflowResponse, error)
	GetWorkflowsByUserID(ctx context.Context, req *dto.GetWorkflowsRequest) (*dto.PaginatedWorkflowsResponse, error)
	GetWorkflowByID(ctx context.Context, req *dto.GetWorkflowByIDRequest) (*dto.WorkflowResponse, error)
	UpdateWorkflow(ctx context.Context, req *dto.UpdateWorkflowRequest) (*dto.WorkflowResponse, error)
	DeleteWorkflow(ctx context.Context, req *dto.DeleteWorkflowRequest) error
}

type workflowService struct {
	queries *repo.Queries
}

func NewWorkflowService(queries *repo.Queries) WorkflowService {
	return &workflowService{
		queries: queries,
	}
}

func (s *workflowService) CreateWorkflow(ctx context.Context, workflow *dto.CreateWorkflowRequest) (*dto.WorkflowResponse, error) {
	newWorkflow, err := s.queries.CreateWorkflow(ctx, repo.CreateWorkflowParams{
		Name:        workflow.Name,
		Description: workflow.Description,
		UserID:      workflow.UserID,
	})
	if err != nil {
		return nil, err
	}
	return toWorkflowResponse(&newWorkflow), nil
}

func (s *workflowService) GetWorkflowsByUserID(ctx context.Context, req *dto.GetWorkflowsRequest) (*dto.PaginatedWorkflowsResponse, error) {
	rows, err := s.queries.GetWorkflowsByUserID(ctx, repo.GetWorkflowsByUserIDParams{
		UserID:  req.UserID,
		Column2: req.Search,
		Limit:   req.Limit,
		Offset:  req.Offset,
	})
	if err != nil {
		return nil, err
	}

	var totalCount int32
	if len(rows) > 0 {
		totalCount = int32(rows[0].TotalCount)
	}
	workflows := make([]repo.Workflow, 0, len(rows))
	for _, row := range rows {
		workflows = append(workflows, repo.Workflow{
			ID:          row.ID,
			UserID:      row.UserID,
			Name:        row.Name,
			Description: row.Description,
			CreatedAt:   row.CreatedAt,
		})
	}

	currentPage := (req.Offset / req.Limit) + 1
	totalPages := (totalCount + req.Limit - 1) / req.Limit

	return &dto.PaginatedWorkflowsResponse{
		Workflows:       workflows,
		HasNextPage:     currentPage < totalPages,
		HasPreviousPage: currentPage > 1,
		TotalCount:      totalCount,
		CurrentPage:     currentPage,
		TotalPages:      totalPages,
	}, nil
}

func (s *workflowService) GetWorkflowByID(ctx context.Context, req *dto.GetWorkflowByIDRequest) (*dto.WorkflowResponse, error) {
	workflow, err := s.queries.GetWorkflowByID(ctx, repo.GetWorkflowByIDParams{
		ID:     req.ID,
		UserID: req.UserID,
	})
	if err != nil {
		return nil, err
	}
	return toWorkflowResponse(&workflow), nil
}

func (s *workflowService) UpdateWorkflow(ctx context.Context, req *dto.UpdateWorkflowRequest) (*dto.WorkflowResponse, error) {
	currentWorkflow, err := s.queries.GetWorkflowByID(ctx, repo.GetWorkflowByIDParams{
		ID:     req.ID,
		UserID: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	name := currentWorkflow.Name
	if req.Name != nil && *req.Name != "" {
		name = *req.Name
	}

	description := currentWorkflow.Description
	if req.Description != nil && *req.Description != "" {
		description = *req.Description
	}

	updatedWorkflow, err := s.queries.UpdateWorkflow(ctx, repo.UpdateWorkflowParams{
		ID:          req.ID,
		Name:        name,
		Description: description,
	})
	if err != nil {
		return nil, err
	}
	return toWorkflowResponse(&updatedWorkflow), nil
}

func (s *workflowService) DeleteWorkflow(ctx context.Context, req *dto.DeleteWorkflowRequest) error {
	_, err := s.queries.GetWorkflowByID(ctx, repo.GetWorkflowByIDParams{
		ID:     req.ID,
		UserID: req.UserID,
	})
	if err != nil {
		return err
	}
	err = s.queries.DeleteWorkflow(ctx, req.ID)
	if err != nil {
		return err
	}
	return nil
}

func toWorkflowResponse(w *repo.Workflow) *dto.WorkflowResponse {
	return &dto.WorkflowResponse{
		ID:          w.ID,
		UserID:      w.UserID,
		Name:        w.Name,
		Description: w.Description,
		CreatedAt:   w.CreatedAt.Format(time.RFC3339),
	}
}
