-- name: CreateConsumable :one
INSERT INTO consumable (
  "createdAt",
  title,
  uom,
  remaining
) VALUES (
  $1,
  $2,
  $3,
  $4
) RETURNING *;