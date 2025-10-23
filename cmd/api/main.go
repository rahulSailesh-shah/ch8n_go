package main

import (
	"context"

	"github.com/rahulSailesh-shah/ch8n_go/internal/app"
	"github.com/rahulSailesh-shah/ch8n_go/internal/server"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	app, err := app.NewApp(ctx, config)
	if err != nil {
		panic(err)
	}

	server := server.NewServer(app, ctx)
	if err := server.Run(); err != nil {
		app.DB.Close()
		panic(err)
	}
}
