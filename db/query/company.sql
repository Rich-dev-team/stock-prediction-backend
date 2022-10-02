
-- name: CreateCompany :one
INSERT INTO company(
    company_name, stock_symbol
) VALUES(
    $1, $2
) RETURNING *;