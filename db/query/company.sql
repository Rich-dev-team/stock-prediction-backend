
-- name: CreateCompany :one
INSERT INTO company(
    company_name, stock_symbol
) VALUES(
    $1, $2
) RETURNING *;

-- name: GetCompany :one
SELECT * FROM company
WHERE company_name = $1 LIMIT 1;

-- name: ListCompany :many
SELECT * FROM company
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateCompanyName :one
UPDATE company 
SET company_name = $2
where id = $1
RETURNING *;

-- name: UpdateCompanyStockSymbol :one
UPDATE company 
SET stock_symbol = $2
where id = $1
RETURNING *;

-- name: DeleteCompany :exec
DELETE FROM company WHERE id = $1;
