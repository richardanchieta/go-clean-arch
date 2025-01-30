-- name: CreateOrder :one
INSERT INTO orders (user_id, product_id, quantity, total)
VALUES ($1, $2, $3, $4)
RETURNING *;
