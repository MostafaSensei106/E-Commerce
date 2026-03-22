-- name: GetAllProducts :many
SELECT *
FROM
 products; 

-- FindProductByID :one
SELECT *
FROM
 products
WHERE
 id = $1;