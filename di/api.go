package di

import (
	"database/sql"
	"github.com/linzhengen/mii-go/config"
	"github.com/linzhengen/mii-go/internal/infrastructure/persistence/mysql/sqlc"
	"github.com/linzhengen/mii-go/internal/infrastructure/trans"
	"github.com/linzhengen/mii-go/internal/infrastructure/user"
	"github.com/linzhengen/mii-go/internal/interface/api/handler"
	"github.com/linzhengen/mii-go/internal/interface/api/router"
	"github.com/linzhengen/mii-go/internal/usecase"
	"go.uber.org/dig"
)

func NewApi(envCfg config.EnvConfig, db *sql.DB) *dig.Container {
	c := dig.New()
	// config
	must(c.Provide(func() config.MySQL {
		return envCfg.MySQL
	}))
	// db
	must(c.Provide(func() *sql.DB {
		return db
	}))
	must(c.Provide(func() sqlc.DBTX {
		return db
	}))

	// domain

	// infrastructure
	must(c.Provide(user.New))
	must(c.Provide(sqlc.New))
	must(c.Provide(trans.New))

	// usecase
	must(c.Provide(usecase.NewUserUseCase))

	// interface
	must(c.Provide(handler.NewHealthHandler))
	must(c.Provide(handler.NewUserHandler))
	must(c.Provide(router.New))

	return c
}
