package mysql

import (
	"context"
	"database/sql"

	"github.com/linzhengen/mii-go/internal/domain/contextx"

	"github.com/linzhengen/mii-go/config"
	"github.com/linzhengen/mii-go/internal/infrastructure/persistence/mysql/sqlc"
)

func NewConn(cfg config.MySQL) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(cfg.MaxLifetime)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	return db, nil
}

func GetQ(ctx context.Context, q *sqlc.Queries) *sqlc.Queries {
	if t, ok := contextx.FromTrans(ctx); ok {
		return t.(*sqlc.Queries)
	}
	return q
}
