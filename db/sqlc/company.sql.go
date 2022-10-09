// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: company.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createCompany = `-- name: CreateCompany :one
INSERT INTO company(
    company_name, stock_symbol
) VALUES(
    $1, $2
) RETURNING id, company_name, stock_symbol, created_at
`

type CreateCompanyParams struct {
	CompanyName string `json:"company_name"`
	StockSymbol string `json:"stock_symbol"`
}

func (q *Queries) CreateCompany(ctx context.Context, arg CreateCompanyParams) (Company, error) {
	row := q.db.QueryRowContext(ctx, createCompany, arg.CompanyName, arg.StockSymbol)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.CompanyName,
		&i.StockSymbol,
		&i.CreatedAt,
	)
	return i, err
}

const deleteCompany = `-- name: DeleteCompany :exec
DELETE FROM company WHERE id = $1
`

func (q *Queries) DeleteCompany(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteCompany, id)
	return err
}

const getCompany = `-- name: GetCompany :one
SELECT id, company_name, stock_symbol, created_at FROM company
WHERE company_name = $1 LIMIT 1
`

func (q *Queries) GetCompany(ctx context.Context, companyName string) (Company, error) {
	row := q.db.QueryRowContext(ctx, getCompany, companyName)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.CompanyName,
		&i.StockSymbol,
		&i.CreatedAt,
	)
	return i, err
}

const listCompany = `-- name: ListCompany :many
SELECT id, company_name, stock_symbol, created_at FROM company
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListCompanyParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCompany(ctx context.Context, arg ListCompanyParams) ([]Company, error) {
	rows, err := q.db.QueryContext(ctx, listCompany, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Company
	for rows.Next() {
		var i Company
		if err := rows.Scan(
			&i.ID,
			&i.CompanyName,
			&i.StockSymbol,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCompanyName = `-- name: UpdateCompanyName :one
UPDATE company 
SET company_name = $2
where id = $1
RETURNING id, company_name, stock_symbol, created_at
`

type UpdateCompanyNameParams struct {
	ID          uuid.UUID `json:"id"`
	CompanyName string    `json:"company_name"`
}

func (q *Queries) UpdateCompanyName(ctx context.Context, arg UpdateCompanyNameParams) (Company, error) {
	row := q.db.QueryRowContext(ctx, updateCompanyName, arg.ID, arg.CompanyName)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.CompanyName,
		&i.StockSymbol,
		&i.CreatedAt,
	)
	return i, err
}

const updateCompanyStockSymbol = `-- name: UpdateCompanyStockSymbol :one
UPDATE company 
SET stock_symbol = $2
where id = $1
RETURNING id, company_name, stock_symbol, created_at
`

type UpdateCompanyStockSymbolParams struct {
	ID          uuid.UUID `json:"id"`
	StockSymbol string    `json:"stock_symbol"`
}

func (q *Queries) UpdateCompanyStockSymbol(ctx context.Context, arg UpdateCompanyStockSymbolParams) (Company, error) {
	row := q.db.QueryRowContext(ctx, updateCompanyStockSymbol, arg.ID, arg.StockSymbol)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.CompanyName,
		&i.StockSymbol,
		&i.CreatedAt,
	)
	return i, err
}
