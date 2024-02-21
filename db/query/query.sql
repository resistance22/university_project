-- name: GetUser :one
SELECT * FROM app_user
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO app_user (
  first_name,   
  last_name,
  email,
  password, 
  ) VALUES ( $1, $2, $3, $4 ) RETURNING *;