package transaction

import (
	"context"
	"database/sql"
)

type tx struct {
	db *sql.DB
}

func NewTransaction(db *sql.DB) Transaction {
	return &tx{db}
}

func (t *tx) DoInTx(ctx context.Context, txKey any, f func(context.Context) (any, error)) (any, error) {
	tx, err := t.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, txKey, tx)
	result, err := f(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}
	return result, nil
}
