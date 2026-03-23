-- name: GetAllProducts :many
SELECT *
FROM products
ORDER BY created_at DESC;


-- name: GetProductByID :one
SELECT *
FROM products
WHERE id = $1;


-- name: CreateProduct :one
INSERT INTO products (
    name,
    price_in_cents,
    quantity
)
VALUES ($1, $2, $3)
RETURNING *;


-- name: UpdateProductWhereID :one
UPDATE products
SET 
    name = $1,
    price_in_cents = $2,
    quantity = $3,
    updated_at = now()
WHERE id = $4
RETURNING *;


-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;


-- name: SearchProducts :many
SELECT *
FROM products
WHERE name ILIKE '%' || $1 || '%'
ORDER BY created_at DESC;


-- name: GetAvailableProducts :many
SELECT *
FROM products
WHERE quantity > 0
ORDER BY created_at DESC;


-- name: IncreaseProductQuantity :exec
UPDATE products
SET 
    quantity = quantity + $1,
    updated_at = now()
WHERE id = $2;



-- name: UpdateProductPrice :exec
UPDATE products
SET 
    price_in_cents = $1,
    updated_at = now()
WHERE id = $2;


-- name: ProductExists :one
SELECT EXISTS (
    SELECT 1 FROM products WHERE id = $1
);



-- name: GetTopProductsByQuantity :many
SELECT *
FROM products
ORDER BY quantity DESC
LIMIT $1;


-- name: CreateOrder :one
INSERT INTO orders (
    customer_id,
    status
) VALUES ($1, $2)
RETURNING *;


-- name: CreateOrderItem :one
INSERT INTO orders_items (
    order_id,
     product_id, quantity , price_in_cents
) VALUES ($1, $2, $3, $4)
RETURNING *;


-- name: GetOrderByID :one
SELECT *
FROM orders
WHERE id = $1;


-- name: GetAllOrders :many
SELECT *
FROM orders
ORDER BY created_at DESC;