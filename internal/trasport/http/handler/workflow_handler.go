package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahulSailesh-shah/ch8n_go/internal/db/repo"
	"github.com/rahulSailesh-shah/ch8n_go/internal/dto"
	"github.com/rahulSailesh-shah/ch8n_go/internal/service"
)

type WorkflowHandler struct {
	workflowService service.WorkflowService
}

func NewWorkflowHandler(workflowService service.WorkflowService) *WorkflowHandler {
	return &WorkflowHandler{
		workflowService: workflowService,
	}
}

func (h *WorkflowHandler) CreateWorkflow(c *gin.Context) {
	var req dto.CreateWorkflowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	workflow, err := h.workflowService.CreateWorkflow(c.Request.Context(),
		&repo.Workflow{
			Name:        req.Name,
			Description: req.Description,
			UserID:      c.MustGet("user_id").(string),
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Failed to create workflow",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Workflow created successfully",
		Data:    workflow,
	})
}

func (h *WorkflowHandler) GetWorkflows(c *gin.Context) {
	workflows, err := h.workflowService.GetWorkflowsByUserID(c.Request.Context(), c.MustGet("user_id").(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Failed to retrieve workflows",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Workflows retrieved successfully",
		Data:    workflows,
	})
}
