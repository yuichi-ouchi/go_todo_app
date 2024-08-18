package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/yuichi-ouchi/go_todo_app/config"
)

func New(ctx context.Context, cfg *config.Config) (*sqlx.DB, func(), error) {
	//"sqlx connect" execute ping command.
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))
	if err != nil {
		return nil, nil, err
	}

	// database open but not connection.
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	// check database alive
	if err := db.PingContext(ctx); err != nil {
		return nil, func() { _ = db.Close() }, err
	}

	xdb := sqlx.NewDb(db, "mysql")
	return xdb, func() { _ = db.Close() }, nil

}
