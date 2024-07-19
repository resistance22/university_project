-- name: CreateConsumable :one
INSERT INTO consumable (
  id,
  "createdAt",
  title,
  uom,
  remaining
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
) RETURNING *;

-- name: GetAllConsumable :many
SELECT * FROM consumable;

-- name: CreateUser :one
INSERT INTO app_user (
  id,
  created_at,
  first_name,
  last_name,
  user_name,
  password
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
) RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM app_user;

-- name: GetUserByUserName :one

SELECT * FROM app_user WHERE user_name = $1 LIMIT 1;