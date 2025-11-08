package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/rahulSailesh-shah/ch8n_go/internal/db/repo"
	"github.com/rahulSailesh-shah/ch8n_go/internal/dto"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/inngest"
)

type WorkflowService interface {
	CreateWorkflow(ctx context.Context, workflow *dto.CreateWorkflowRequest) (*dto.WorkflowResponse, error)
	GetWorkflowsByUserID(ctx context.Context, req *dto.GetWorkflowsRequest) (*dto.PaginatedWorkflowsResponse, error)
	GetWorkflowByID(ctx context.Context, req *dto.GetWorkflowByIDRequest) (*dto.WorkflowResponse, error)
	UpdateWorkflowName(ctx context.Context, req *dto.UpdateWorkflowNameRequest) (*dto.WorkflowResponse, error)
	UpdateWorkflow(ctx context.Context, req *dto.UpdateWorkflowRequest) (*dto.WorkflowResponse, error)
	DeleteWorkflow(ctx context.Context, req *dto.DeleteWorkflowRequest) error
}

type workflowService struct {
	queries *repo.Queries
	inngest *inngest.Inngest
	db      *pgxpool.Pool
}

func NewWorkflowService(queries *repo.Queries, db *pgxpool.Pool, inngest *inngest.Inngest) WorkflowService {
	return &workflowService{
		queries: queries,
		db:      db,
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
	node, err := s.queries.CreateNode(ctx, repo.CreateNodeParams{
		ID:         uuid.New(),
		WorkflowID: newWorkflow.ID,
		Name:       string(dto.NodeTypeInitial),
		Type:       string(dto.NodeTypeInitial),
		Position:   []byte(`{"x": 0, "y": 0}`),
	})
	if err != nil {
		return nil, err
	}
	return toWorkflowResponse(&newWorkflow, []repo.Node{node}, []repo.Connection{}), nil
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
			UpdatedAt:   row.UpdatedAt,
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

	nodes, err := s.queries.GetNodesByWorkflowID(ctx, workflow.ID)
	if err != nil {
		return nil, err
	}

	edges, err := s.queries.GetConnectionsByWorkflowID(ctx, workflow.ID)
	if err != nil {
		return nil, err
	}
	return toWorkflowResponse(&workflow, nodes, edges), nil
}

func (s *workflowService) UpdateWorkflowName(ctx context.Context, req *dto.UpdateWorkflowNameRequest) (*dto.WorkflowResponse, error) {
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

	updatedWorkflow, err := s.queries.UpdateWorkflowName(ctx, repo.UpdateWorkflowNameParams{
		ID:   req.ID,
		Name: name,
	})
	if err != nil {
		return nil, err
	}

	nodes, err := s.queries.GetNodesByWorkflowID(ctx, updatedWorkflow.ID)
	if err != nil {
		return nil, err
	}

	edges, err := s.queries.GetConnectionsByWorkflowID(ctx, updatedWorkflow.ID)
	if err != nil {
		return nil, err
	}

	return toWorkflowResponse(&updatedWorkflow, nodes, edges), nil
}

func (s *workflowService) DeleteWorkflow(ctx context.Context, req *dto.DeleteWorkflowRequest) error {
	_, err := s.queries.GetWorkflowByID(ctx, repo.GetWorkflowByIDParams{
		ID:     req.ID,
		UserID: req.UserID,
	})
	if err != nil {
		return err
	}
	// Run everything in a transaction
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	q := repo.New(tx)

	// Delete connections first (foreign key constraint)
	if err := q.DeleteConnectionsByWorkflowID(ctx, req.ID); err != nil {
		return fmt.Errorf("failed to delete connections: %w", err)
	}

	// Delete nodes next (foreign key constraint)
	if err := q.DeleteNodesByWorkflowID(ctx, req.ID); err != nil {
		return fmt.Errorf("failed to delete nodes: %w", err)
	}

	err = s.queries.DeleteWorkflow(ctx, req.ID)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *workflowService) UpdateWorkflow(ctx context.Context, req *dto.UpdateWorkflowRequest) (*dto.WorkflowResponse, error) {
	workflow, err := s.queries.GetWorkflowByID(ctx, repo.GetWorkflowByIDParams{
		ID:     req.ID,
		UserID: req.UserID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get workflow: %w", err)
	}

	// Run everything in a transaction
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	q := repo.New(tx)

	// Delete connections first (foreign key constraint)
	if err := q.DeleteConnectionsByWorkflowID(ctx, workflow.ID); err != nil {
		return nil, fmt.Errorf("failed to delete connections: %w", err)
	}

	// Delete nodes next (foreign key constraint)
	if err := q.DeleteNodesByWorkflowID(ctx, workflow.ID); err != nil {
		return nil, fmt.Errorf("failed to delete nodes: %w", err)
	}

	// Create new nodes
	newNodeParams, err := toCreateNodeParams(req.Nodes, workflow.ID)
	if err != nil {
		return nil, err
	}

	createdNodes := make([]repo.Node, 0, len(newNodeParams))
	for _, nodeParam := range newNodeParams {
		node, err := q.CreateNode(ctx, nodeParam)
		if err != nil {
			return nil, fmt.Errorf("failed to create node %s: %w", nodeParam.Name, err)
		}
		createdNodes = append(createdNodes, node)
	}

	// Create new connections
	newEdgeParams, err := toCreateConnectionParams(req.Edges, workflow.ID)
	if err != nil {
		return nil, err
	}

	createdEdges := make([]repo.Connection, 0, len(newEdgeParams))
	for _, edgeParam := range newEdgeParams {
		edge, err := q.CreateConnection(ctx, edgeParam)
		if err != nil {
			return nil, fmt.Errorf("failed to create connection: %w", err)
		}
		createdEdges = append(createdEdges, edge)
	}

	// Touch workflow to update updated_at timestamp
	updatedWorkflow, err := q.UpdateWorkflowName(ctx, repo.UpdateWorkflowNameParams{
		ID:   workflow.ID,
		Name: workflow.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update workflow timestamp: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return toWorkflowResponse(&updatedWorkflow, createdNodes, createdEdges), nil
}

func toCreateConnectionParams(edges []dto.UpdateConnectionRequest, workflowID uuid.UUID) ([]repo.CreateConnectionParams, error) {
	createEdges := make([]repo.CreateConnectionParams, 0, len(edges))
	for _, edge := range edges {
		createEdges = append(createEdges, repo.CreateConnectionParams{
			WorkflowID:   workflowID,
			SourceNodeID: edge.Source,
			TargetNodeID: edge.Target,
			FromOutput:   edge.SourceHandle,
			ToInput:      edge.TargetHandle,
		})
	}
	return createEdges, nil
}

func toCreateNodeParams(nodes []dto.UpdateNodeRequest, workflowID uuid.UUID) ([]repo.CreateNodeParams, error) {
	createNodes := make([]repo.CreateNodeParams, 0, len(nodes))
	for _, node := range nodes {
		positionJSON, err := json.Marshal(node.Position)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal position for node %s: %w", node.Name, err)
		}

		// Ensure Data is not nil
		data := node.Data
		if data == nil {
			data = []byte(`{}`)
		}

		createNodes = append(createNodes, repo.CreateNodeParams{
			ID:         node.ID,
			WorkflowID: workflowID,
			Name:       node.Name,
			Type:       string(node.Type),
			Position:   positionJSON,
			Data:       data,
		})
	}
	return createNodes, nil
}

func toWorkflowResponse(w *repo.Workflow, nodes []repo.Node, edges []repo.Connection) *dto.WorkflowResponse {
	return &dto.WorkflowResponse{
		ID:          w.ID,
		UserID:      w.UserID,
		Name:        w.Name,
		Description: w.Description,
		CreatedAt:   w.CreatedAt,
		UpdatedAt:   w.UpdatedAt,
		Nodes:       nodes,
		Edges:       edges,
	}
}
