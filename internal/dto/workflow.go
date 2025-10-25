package dto

import (
	"github.com/rahulSailesh-shah/ch8n_go/internal/db/repo"
)

type CreateWorkflowRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
}

type UpdateWorkflowRequest struct {
	UserID      string  `json:"user_id" validate:"required"`
	ID          int32   `json:"id" validate:"required"`
	Name        *string `json:"name,omitempty" validate:"omitempty,min=3,max=100"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=500"`
}

type DeleteWorkflowRequest struct {
	UserID string `json:"user_id" validate:"required"`
	ID     int32  `json:"id" validate:"required"`
}

type GetWorkflowByIDRequest struct {
	UserID string `json:"user_id" validate:"required"`
	ID     int32  `json:"id" validate:"required"`
}

type GetWorkflowsRequest struct {
	Search string `json:"search" binding:"required"`
	Limit  int32  `json:"limit" binding:"required"`
	Offset int32  `json:"offset" binding:"required"`
	UserID string `json:"user_id" binding:"required"`
}

type PaginatedWorkflowsResponse struct {
	Workflows       []repo.Workflow `json:"workflows"`
	HasNextPage     bool            `json:"has_next_page"`
	HasPreviousPage bool            `json:"has_previous_page"`
	TotalCount      int32           `json:"total_count"`
	CurrentPage     int32           `json:"current_page"`
	TotalPages      int32           `json:"total_pages"`
}

type WorkflowResponse struct {
	ID          int32  `json:"id"`
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
