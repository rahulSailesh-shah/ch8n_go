package http

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/rahulSailesh-shah/ch8n_go/internal/service"
	"github.com/rahulSailesh-shah/ch8n_go/internal/trasport/http/handler"
	"github.com/rahulSailesh-shah/ch8n_go/internal/trasport/http/middleware"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/config"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/inngest"
)

func RegisterRoutes(r *gin.Engine, service service.Service, inngest inngest.Inngest, authKeys jwk.Set, polarConfig *config.PolarConfig) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Inngest Endpoint
	r.Any("/api/inngest", inngest.Handler())

	// Webhook Endpoint,
	// TODO: Protect this endpoint with auth token
	webhookHandler := handler.NewWebhookHandler(service.Webhook)
	r.POST("/api/webhook/:workflow_id", webhookHandler.HandleWebhook)

	// Protected routes
	protectedGroup := r.Group("")
	protectedGroup.Use(middleware.AuthMiddleware(authKeys))
	protectedGroup.Use(middleware.SubscriptionMiddleware(polarConfig))

	// Workflow routes
	workflowHandler := handler.NewWorkflowHandler(service.Workflow)
	workflowGroup := protectedGroup.Group("/workflows")
	{
		workflowGroup.POST("", workflowHandler.CreateWorkflow)
		workflowGroup.GET("", workflowHandler.GetWorkflowsByUserID)
		workflowGroup.GET("/:id", workflowHandler.GetWorkflowByID)
		workflowGroup.PUT("/:id/name", workflowHandler.UpdateWorkflowName)
		workflowGroup.PUT("/:id", workflowHandler.UpdateWorkflow)
		workflowGroup.DELETE("/:id", workflowHandler.DeleteWorkflow)
		workflowGroup.POST("/:id/execute", workflowHandler.ExecuteWorkflow)
	}
}
