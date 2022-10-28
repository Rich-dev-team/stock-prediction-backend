package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore to create new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	fnErr := fn(q)

	if fnErr != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return fnErr
	}
	return tx.Commit()
}

func (store *Store) CreateStocksTx(ctx context.Context, arg CreateBatchStockPriceParams) ([]StockPrice, error) {
	var results []StockPrice

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		results, err = q.CreateBatchStockPrice(ctx, CreateBatchStockPriceParams{
			CompanysID: arg.CompanysID,
			Prices:     arg.Prices,
			CreatedAts: arg.CreatedAts,
		})
		if err != nil {
			return err
		}
		return nil
	})

	return results, err
}
