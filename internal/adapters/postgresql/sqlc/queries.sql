-- name: GetAllProducts :many
SELECT *
FROM
 products; 

-- name: GetProductByID :one
SELECT *
FROM
 products
WHERE id = $1;

-- name: CreateProduct :one
INSERT INTO
 products (name, price_in_cents, quantity)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateProduct :exec
UPDATE products
SET name = $1, price_in_cents = $2, quantity = $3
WHERE id = $4;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;