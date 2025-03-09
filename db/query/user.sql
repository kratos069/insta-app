-- name: CreateUser :one
INSERT INTO users (
  username, hashed_password, full_name, email, profile_picture, bio
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;