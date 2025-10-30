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
	"github.com/rahulSailesh-shah/ch8n_go/pkg/auth"
)

type Server struct {
	App        *app.App
	ctx        context.Context
	Engine     *gin.Engine
	httpServer *httpSrv.Server
}

func NewServer(app *app.App, ctx context.Context) *Server {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	authKeys, err := auth.LoadKeys(app.Config.Auth.JwksURL)
	if err != nil {
		panic(err)
	}

	http.RegisterRoutes(engine, *app.Service, authKeys, &app.Config.Polar)

	srv := &httpSrv.Server{
		Addr:    fmt.Sprintf(":%d", app.Config.Server.Port),
		Handler: engine,
	}

	return &Server{
		App:        app,
		ctx:        ctx,
		Engine:     engine,
		httpServer: srv,
	}
}

func (s *Server) Run() error {
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
	ctx, stop := signal.NotifyContext(s.ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	defer s.App.DB.Close()
	stop()

	timeout := time.Duration(s.App.Config.Server.GracefulShutdownSec) * time.Second
	ctx, cancel := context.WithTimeout(s.ctx, timeout)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")
	done <- true
}
