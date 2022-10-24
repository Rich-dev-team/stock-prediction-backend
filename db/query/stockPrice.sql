-- name: CreateStockPrice :one
INSERT INTO stock_price(company_id, price, created_at)
VALUES($1, $2, $3)
RETURNING *;
-- name: CreateBatchStockPrice :many
INSERT INTO stock_price(company_id, price, created_at)
VALUES (
        UNNEST (@companys_id::bigint []),
        UNNEST (@prices::int []),
        UNNEST (@created_ats::timestamptz [])
    )
RETURNING *;
-- name: UpdateStockPrice :one
UPDATE stock_price
SET price = $2
where id = $1
RETURNING *;
-- name: ListStockPriceByRange :many
SELECT *
FROM stock_price
WHERE company_id = $1
    AND created_at BETWEEN sqlc.arg(startTime) AND sqlc.arg(endTime)
ORDER BY created_at
LIMIT $2;
-- name: ListStockPriceByRangeForUpdate :many
SELECT *
FROM stock_price
WHERE company_id = $1
    AND created_at BETWEEN sqlc.arg(startTime) AND sqlc.arg(endTime)
ORDER BY created_at
LIMIT $2
For NO KEY UPDATE;
-- name: ListAllStockPrice :many
SELECT *
FROM stock_price
where company_id = $1
ORDER BY created_at
LIMIT $2 OFFSET $3;
-- name: DeleteStockPrice :exec
DELETE FROM stock_price
WHERE id = $1;