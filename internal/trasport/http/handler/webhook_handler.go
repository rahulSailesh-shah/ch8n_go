package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rahulSailesh-shah/ch8n_go/internal/dto"
	"github.com/rahulSailesh-shah/ch8n_go/internal/service"
)

type WebhookHandler struct {
	webhookService service.WebhookService
}

func NewWebhookHandler(webhookService service.WebhookService,
) *WebhookHandler {
	return &WebhookHandler{
		webhookService: webhookService,
	}
}

func (h *WebhookHandler) HandleWebhook(c *gin.Context) {
	workflowID, err := uuid.Parse(c.Param("workflow_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid workflow ID",
			Error:   err.Error(),
		})
		return
	}

	workflow, err := h.webhookService.ProcessWebhook(c.Request.Context(), &dto.WebhookRequest{
		WorkflowID: workflowID,
		Data:       c.Request.Body,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Failed to handle webhook",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Workflow executed successfully",
		Data:    workflow,
	})
}
