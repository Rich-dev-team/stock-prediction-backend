-- name: CreateStockPrice :one
INSERT INTO stock_price(company_id, price, created_at)
VALUES($1, $2, $3)
RETURNING *;
-- name: UpdateStockPrice :one
UPDATE stock_price
SET price = $2
where id = $1
RETURNING *;
-- name: ListStockPriceByRange :many
SELECT *
FROM stock_price
WHERE company_id = $1 AND created_at BETWEEN $2 AND $3
ORDER BY created_at
LIMIT $4;
-- name: ListAllStockPrice :many
SELECT *
FROM stock_price
where company_id = $1
ORDER BY created_at
LIMIT $2 OFFSET $3;
-- name: DeleteStockPrice :exec
DELETE FROM stock_price
WHERE id = $1;