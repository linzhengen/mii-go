package trans

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/linzhengen/mii-go/app/domain/contextx"
	"github.com/linzhengen/mii-go/app/domain/trans"
)

func New(db *sql.DB) trans.Repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *sql.DB
}

func (a *repository) ExecTrans(ctx context.Context, fn func(context.Context) error) error {
	if _, ok := contextx.FromTrans(ctx); ok {
		return fn(ctx)
	}

	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	err = fn(contextx.NewTrans(ctx, tx))
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
