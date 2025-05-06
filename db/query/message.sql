

-- name: CreateOrders :one
INSERT INTO "orders" (customer_id, product_id, order_status, total_amount)
VALUES ($1, $2, 'PENDING', $3)
RETURNING *;

-- name: UpdateOrdersByID :one
UPDATE "orders" 
SET id = $1, order_status = $2
RETURNING *;

-- name: CreateCustomer :many
INSERT INTO "customer" (customer_name, contact)
VALUES ($1, $2)
RETURNING *;
-- name: CreateProduct :many
INSERT INTO "product" (product_name, price)
VALUES ($1, $2)
RETURNING *;

-- name: GetProductByID :many
SELECT * FROM product
WHERE id = $1;  

-- name: GetCustomerByID :one
SELECT * FROM "customer"
WHERE id = $1;  