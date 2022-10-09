-- name: CreateStockPrice :one
INSERT INTO stock_price(company_id, price)
VALUES($1, $2)
RETURNING *;
-- name: ListStockPriceByRange :many
SELECT *
FROM stock_price
WHERE company_id = $1 BETWEEN $2 AND $3
ORDER BY cur_date
LIMIT $4;
-- name: ListAllStockPrice :many
SELECT *
FROM stock_price
where company_id = $1
ORDER BY cur_date
LIMIT $2 OFFSET $3;