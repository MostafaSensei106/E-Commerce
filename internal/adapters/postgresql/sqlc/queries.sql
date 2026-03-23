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


-- name: CreateOrderItem :one
INSERT INTO orders_items (
    order_id,
    product_id,
    quantity,
    price_in_cents
)
VALUES ($1, $2, $3, $4)
RETURNING *;


-- name: GetOrderItemsByOrderID :many
SELECT *
FROM orders_items
WHERE order_id = $1;


-- name: GetOrderItem :one
SELECT *
FROM orders_items
WHERE order_id = $1 AND product_id = $2;


-- name: UpdateOrderItem :exec
UPDATE orders_items
SET quantity = $1,
    price_in_cents = $2,
    updated_at = now()
WHERE order_id = $3 AND product_id = $4;


-- name: DeleteOrderItem :exec
DELETE FROM orders_items
WHERE order_id = $1 AND product_id = $2;



-- name: GetOrderWithItems :many
SELECT 
    o.id AS order_id,
    o.status,

    p.id AS product_id,
    p.name,

    oi.quantity,
    oi.price_in_cents

FROM orders o
JOIN orders_items oi ON oi.order_id = o.id
JOIN products p ON p.id = oi.product_id
WHERE o.id = $1;


-- name: GetOrderTotal :one
SELECT 
    COALESCE(SUM(oi.quantity * oi.price_in_cents), 0) AS total
FROM orders_items oi
WHERE oi.order_id = $1;


-- name: DecreaseProductQuantity :exec
UPDATE products
SET quantity = quantity - $1
WHERE id = $2 AND quantity >= $1;