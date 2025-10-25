package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	cfg "github.com/rahulSailesh-shah/ch8n_go/pkg/config"
)

type DB interface {
	Close() error
	Connect() error
	GetDB() *pgx.Conn
}

type db struct {
	ctx context.Context
	db  *pgx.Conn
	cfg cfg.DBConfig
}

func NewDB(ctx context.Context, cfg cfg.DBConfig) DB {
	return &db{
		ctx: ctx,
		db:  nil,
		cfg: cfg,
	}
}

func (d *db) GetDB() *pgx.Conn {
	return d.db
}

func (d *db) Close() error {
	log.Println("Closing database connection")
	if d.db == nil {
		return fmt.Errorf("database connection is already closed")
	}
	return d.db.Close(d.ctx)
}

func (d *db) Connect() error {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		d.cfg.User, d.cfg.Password, d.cfg.Host, d.cfg.Port, d.cfg.Name)

	db, err := pgx.Connect(d.ctx, dsn)
	if err != nil {
		return err
	}

	d.db = db
	log.Println("Connected to Database")
	return nil
}
