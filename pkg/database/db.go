package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
	cfg "github.com/rahulSailesh-shah/ch8n_go/pkg/config"
)

type DB interface {
	Close() error
	Connect() error
	GetDB() *sqlx.DB
}

type db struct {
	ctx context.Context
	db  *sqlx.DB
	cfg cfg.DBConfig
}

func NewDB(ctx context.Context, cfg cfg.DBConfig) DB {
	return &db{
		ctx: ctx,
		db:  nil,
		cfg: cfg,
	}
}

func (d *db) GetDB() *sqlx.DB {
	return d.db
}

func (d *db) Close() error {
	log.Println("Closing database connection")
	if d.db == nil {
		return fmt.Errorf("database connection is already closed")
	}
	return d.db.Close()
}

func (d *db) Connect() error {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		d.cfg.User, d.cfg.Password, d.cfg.Host, d.cfg.Port, d.cfg.Name)

	db, err := sqlx.ConnectContext(d.ctx, d.cfg.Driver, dsn)
	if err != nil {
		return err
	}

	d.db = db
	log.Println("Connected to Database")
	return nil
}
