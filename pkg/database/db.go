package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	cfg "github.com/rahulSailesh-shah/ch8n_go/pkg/config"
)

type DB interface {
	Close() error
	Connect() error
	GetDB() *pgxpool.Pool
}

type db struct {
	ctx context.Context
	db  *pgxpool.Pool
	cfg cfg.DBConfig
}

func NewDB(ctx context.Context, cfg cfg.DBConfig) DB {
	return &db{
		ctx: ctx,
		db:  nil,
		cfg: cfg,
	}
}

func (d *db) GetDB() *pgxpool.Pool {
	return d.db
}

func (d *db) Close() error {
	log.Println("Closing database connection")
	if d.db == nil {
		return fmt.Errorf("database connection is already closed")
	}
	d.db.Close()
	return nil
}

func (d *db) Connect() error {
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		return err
	}

	config.ConnConfig.Host = d.cfg.Host
	config.ConnConfig.Port = uint16(d.cfg.Port)
	config.ConnConfig.User = d.cfg.User
	config.ConnConfig.Password = d.cfg.Password
	config.ConnConfig.Database = d.cfg.Name

	config.MaxConns = 50
	config.MinConns = 10
	config.MaxConnLifetime = 1 * time.Hour
	config.MaxConnIdleTime = 30 * time.Minute
	config.HealthCheckPeriod = 1 * time.Minute

	db, err := pgxpool.NewWithConfig(d.ctx, config)
	if err != nil {
		return err
	}

	d.db = db
	log.Println("Connected to Database with connection pool")
	return nil
}
