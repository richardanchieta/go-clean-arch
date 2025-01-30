-- name: CreateProduct :one
INSERT INTO products (name, description, price)
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListProducts :many
SELECT * FROM products ORDER BY created_at DESC;
