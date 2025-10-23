package dto

import "time"

type CreateWorkflowRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateWorkflowRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type WorkflowResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserID      string    `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type WorkflowListResponse struct {
	Workflows []WorkflowResponse `json:"workflows"`
}
