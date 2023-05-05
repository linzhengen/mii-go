package di

import (
	"database/sql"
	"log"

	"github.com/linzhengen/mii-go/internal/interface/grpc/register"

	"github.com/linzhengen/mii-go/config"
	"github.com/linzhengen/mii-go/internal/infrastructure/persistence/mysql/sqlc"
	"github.com/linzhengen/mii-go/internal/infrastructure/trans"
	"github.com/linzhengen/mii-go/internal/infrastructure/user"
	grpcHandler "github.com/linzhengen/mii-go/internal/interface/grpc/handler"
	"github.com/linzhengen/mii-go/internal/interface/rest/handler"
	"github.com/linzhengen/mii-go/internal/interface/rest/router"
	"github.com/linzhengen/mii-go/internal/usecase"
	"go.uber.org/dig"
)

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func NewDI(envCfg config.EnvConfig, db *sql.DB) *dig.Container {
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

	// interface (rest)
	must(c.Provide(handler.NewHealthHandler))
	must(c.Provide(handler.NewUserHandler))
	must(c.Provide(router.New))

	// interface (grpc)
	must(c.Provide(grpcHandler.NewUserHandler))
	must(c.Provide(register.New))
	return c
}
