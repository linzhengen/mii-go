package di

import (
	"database/sql"
	"github.com/linzhengen/mii-go/app/infrastructure/persistence/mysql/sqlc"
	"github.com/linzhengen/mii-go/app/infrastructure/trans"
	"github.com/linzhengen/mii-go/app/infrastructure/user"
	"github.com/linzhengen/mii-go/app/interface/api/handler"
	"github.com/linzhengen/mii-go/app/interface/api/router"
	"github.com/linzhengen/mii-go/app/usecase"
	"github.com/linzhengen/mii-go/config"
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
