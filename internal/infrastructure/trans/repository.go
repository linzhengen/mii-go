package trans

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/linzhengen/mii-go/internal/infrastructure/persistence/mysql/sqlc"

	"github.com/linzhengen/mii-go/internal/domain/contextx"
	"github.com/linzhengen/mii-go/internal/domain/trans"
)

func New(db *sql.DB, q *sqlc.Queries) trans.Repository {
	return &repository{
		db: db,
		q:  q,
	}
}

type repository struct {
	db *sql.DB
	q  *sqlc.Queries
}

func (a *repository) ExecTrans(ctx context.Context, fn func(context.Context) error) (txErr error) {
	if _, ok := contextx.FromTrans(ctx); ok {
		return fn(ctx)
	}

	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			txErr = fmt.Errorf("rb err: %w", err)
			return
		}
	}()

	qTx := a.q.WithTx(tx)
	if err := fn(contextx.NewTrans(ctx, qTx)); err != nil {
		return err
	}
	return tx.Commit()
}
