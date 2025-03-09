package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store can do both queries and transcations
type Store interface {
	Querier
}

// SQLStore provides all funcs for SQL queries and transactions
type SQLStore struct {
	// queries only supports queries not transactions,
	// so we use it in store struct and add more functionality.
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// executes a function within a DB Transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("txErr: %v, rbErr: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}