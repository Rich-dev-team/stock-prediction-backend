package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Exchange provides all functions to execute db queries and transactions
type Exchange struct {
	*Queries
	db *sql.DB
}

// NewExchange to create new Store
func NewExchange(db *sql.DB) *Exchange {
	return &Exchange{
		db:      db,
		Queries: New(db),
	}
}

func (exchange *Exchange) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := exchange.db.BeginTx(ctx, nil)
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

func (exchange *Exchange) CreateStocksTx(ctx context.Context, arg CreateBatchStockPriceParams) ([]StockPrice, error) {
	var results []StockPrice

	err := exchange.execTx(ctx, func(q *Queries) error {
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
