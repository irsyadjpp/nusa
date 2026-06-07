package database

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Transaction struct {
	tx *sqlx.Tx
}

func (d *Database) BeginTx(ctx context.Context) (*Transaction, error) {
	tx, err := d.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &Transaction{tx: tx}, nil
}

func (t *Transaction) Commit() error {
	return t.tx.Commit()
}

func (t *Transaction) Rollback() error {
	return t.tx.Rollback()
}

func (t *Transaction) GetTx() *sqlx.Tx {
	return t.tx
}
