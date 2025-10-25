package dto

import (
	"time"

	"github.com/rahulSailesh-shah/ch8n_go/internal/db/repo"
)

type CreateWorkflowRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description,omitempty"`
	UserID      string  `json:"-"`
}

type UpdateWorkflowRequest struct {
	UserID      string  `json:"userId" validate:"required"`
	ID          int32   `json:"id" validate:"required"`
	Name        *string `json:"name,omitempty" validate:"omitempty,min=3,max=100"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=500"`
}

type DeleteWorkflowRequest struct {
	UserID string `json:"userId" validate:"required"`
	ID     int32  `json:"id" validate:"required"`
}

type GetWorkflowByIDRequest struct {
	UserID string `json:"userId" validate:"required"`
	ID     int32  `json:"id" validate:"required"`
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
	ID          int32     `json:"id"`
	UserID      string    `json:"userId"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
