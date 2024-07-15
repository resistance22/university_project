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