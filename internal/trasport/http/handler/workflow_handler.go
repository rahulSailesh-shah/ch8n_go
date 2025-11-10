package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	req.UserID = c.MustGet("user_id").(string)

	fmt.Println(req.UserID)

	workflow, err := h.workflowService.CreateWorkflow(c.Request.Context(), &req)
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

func (h *WorkflowHandler) GetWorkflowsByUserID(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	workflows, err := h.workflowService.GetWorkflowsByUserID(c.Request.Context(), &dto.GetWorkflowsRequest{
		Search: search,
		Limit:  int32(limit),
		Offset: int32((page - 1) * limit),
		UserID: c.MustGet("user_id").(string),
	})
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

func (h *WorkflowHandler) GetWorkflowByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid workflow ID",
			Error:   err.Error(),
		})
		return
	}

	workflow, err := h.workflowService.GetWorkflowByID(c.Request.Context(), &dto.GetWorkflowByIDRequest{
		ID:     id,
		UserID: c.MustGet("user_id").(string),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Failed to retrieve workflow",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Workflow retrieved successfully",
		Data:    workflow,
	})
}

func (h *WorkflowHandler) UpdateWorkflowName(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid workflow ID",
			Error:   err.Error(),
		})
		return
	}

	var req dto.UpdateWorkflowNameRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	req.ID = id
	req.UserID = c.MustGet("user_id").(string)

	workflow, err := h.workflowService.UpdateWorkflowName(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Failed to update workflow",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Workflow updated successfully",
		Data:    workflow,
	})
}

func (h *WorkflowHandler) UpdateWorkflow(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid workflow ID",
			Error:   err.Error(),
		})
		return
	}

	var req dto.UpdateWorkflowRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	req.ID = id
	req.UserID = c.MustGet("user_id").(string)

	workflow, err := h.workflowService.UpdateWorkflow(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Failed to update workflow",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Workflow updated successfully",
		Data:    workflow,
	})
}

func (h *WorkflowHandler) DeleteWorkflow(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid workflow ID",
			Error:   err.Error(),
		})
		return
	}
	err = h.workflowService.DeleteWorkflow(c.Request.Context(), &dto.DeleteWorkflowRequest{
		ID:     id,
		UserID: c.MustGet("user_id").(string),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Failed to delete workflow",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Workflow deleted successfully",
	})
}

func (h *WorkflowHandler) ExecuteWorkflow(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid workflow ID",
			Error:   err.Error(),
		})
		return
	}
	workflow, err := h.workflowService.ExecuteWorkflow(c.Request.Context(), &dto.ExecuteWorkflowRequest{
		ID:     id,
		UserID: c.MustGet("user_id").(string),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Failed to execute workflow",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Workflow executed successfully",
		Data:    workflow,
	})
}
