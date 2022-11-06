// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	CreateBatchStockPrice(ctx context.Context, arg CreateBatchStockPriceParams) ([]StockPrice, error)
	CreateCompany(ctx context.Context, arg CreateCompanyParams) (Company, error)
	CreateStockPrice(ctx context.Context, arg CreateStockPriceParams) (StockPrice, error)
	DeleteCompany(ctx context.Context, id int64) error
	DeleteStockPrice(ctx context.Context, id int64) error
	GetCompanyById(ctx context.Context, id int64) (Company, error)
	GetCompanyByName(ctx context.Context, companyName string) (Company, error)
	ListAllStockPrice(ctx context.Context, arg ListAllStockPriceParams) ([]StockPrice, error)
	ListCompany(ctx context.Context, arg ListCompanyParams) ([]Company, error)
	ListStockPriceByRange(ctx context.Context, arg ListStockPriceByRangeParams) ([]StockPrice, error)
	ListStockPriceByRangeForUpdate(ctx context.Context, arg ListStockPriceByRangeForUpdateParams) ([]StockPrice, error)
	UpdateCompanyName(ctx context.Context, arg UpdateCompanyNameParams) (Company, error)
	UpdateCompanyStockSymbol(ctx context.Context, arg UpdateCompanyStockSymbolParams) (Company, error)
	UpdateStockPrice(ctx context.Context, arg UpdateStockPriceParams) (StockPrice, error)
}

var _ Querier = (*Queries)(nil)