-- name: CreatePost :one
INSERT INTO posts (
  user_id, content_url, caption
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetPostByID :one
SELECT * FROM posts
WHERE post_id = $1 LIMIT 1;

-- name: ListPostsByUser :many
SELECT * FROM posts
WHERE user_id = $1
ORDER BY post_id;

-- name: UpdatePost :one
UPDATE posts
SET caption = $2
WHERE post_id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE post_id = $1;