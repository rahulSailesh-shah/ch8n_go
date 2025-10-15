package server

import (
	"context"
	"fmt"
	"log"
	httpSrv "net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rahulSailesh-shah/ch8n_go/internal/app"
	"github.com/rahulSailesh-shah/ch8n_go/internal/trasport/http"
)

type Server struct {
	App        *app.App
	Engine     *gin.Engine
	httpServer *httpSrv.Server
}

func NewServer(app *app.App) *Server {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	http.RegisterRoutes(engine, *app.Service)

	srv := &httpSrv.Server{
		Addr:    fmt.Sprintf(":%d", app.Config.Server.Port),
		Handler: engine,
	}

	return &Server{
		App:        app,
		Engine:     engine,
		httpServer: srv,
	}
}

func (s *Server) Run(ctx context.Context) error {
	done := make(chan bool, 1)
	go s.gracefulShutdown(done)

	log.Printf("Starting server on port %d", s.App.Config.Server.Port)
	if err := s.httpServer.ListenAndServe(); err != nil && err != httpSrv.ErrServerClosed {
		return fmt.Errorf("could not start server: %w", err)
	}

	<-done
	return nil
}

func (s *Server) gracefulShutdown(done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	stop()

	timeout := time.Duration(s.App.Config.Server.GracefulShutdownSec) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")
	done <- true
}
