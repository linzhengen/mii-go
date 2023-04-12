package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/linzhengen/mii-go/app/domain/contextx"
)

type TransFunc func(ctx context.Context) error

func ExecTrans(ctx context.Context, db *sql.DB, fn TransFunc) error {
	if _, ok := contextx.FromTrans(ctx); ok {
		return fn(ctx)
	}

	tx, err := db.BeginTx(ctx, nil)
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
