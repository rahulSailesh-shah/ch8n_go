package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/rahulSailesh-shah/ch8n_go/internal/db/repo"
)

type CreateWorkflowRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description,omitempty"`
	UserID      string  `json:"-"`
}

type UpdateWorkflowNameRequest struct {
	UserID string    `json:"-"`
	ID     uuid.UUID `json:"-"`
	Name   *string   `json:"name,omitempty" validate:"omitempty,min=3,max=100"`
}

type UpdateWorkflowRequest struct {
	ID     uuid.UUID                 `json:"id" validate:"required"`
	UserID string                    `json:"-"`
	Nodes  []UpdateNodeRequest       `json:"nodes"`
	Edges  []UpdateConnectionRequest `json:"edges"`
}

type DeleteWorkflowRequest struct {
	UserID string    `json:"userId" validate:"required"`
	ID     uuid.UUID `json:"id" validate:"required"`
}

type GetWorkflowByIDRequest struct {
	UserID string    `json:"userId" validate:"required"`
	ID     uuid.UUID `json:"id" validate:"required"`
}

type GetWorkflowsRequest struct {
	Search string `json:"search" binding:"required"`
	Limit  int32  `json:"limit" binding:"required"`
	Offset int32  `json:"offset" binding:"required"`
	UserID string `json:"userId" binding:"required"`
}

type PaginatedWorkflowsResponse struct {
	Workflows       []repo.Workflow `json:"workflows"`
	HasNextPage     bool            `json:"hasNextPage"`
	HasPreviousPage bool            `json:"hasPreviousPage"`
	TotalCount      int32           `json:"totalCount"`
	CurrentPage     int32           `json:"currentPage"`
	TotalPages      int32           `json:"totalPages"`
}

type WorkflowResponse struct {
	ID          uuid.UUID         `json:"id"`
	UserID      string            `json:"userId"`
	Name        string            `json:"name"`
	Description *string           `json:"description"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
	Nodes       []repo.Node       `json:"nodes"`
	Edges       []repo.Connection `json:"edges"`
}

type ExecuteWorkflowRequest struct {
	ID     uuid.UUID `json:"-"`
	UserID string    `json:"-"`
}
