-- name: CreateComment :one
INSERT INTO comments (
  user_id, post_id, content
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetCommentByID :one
SELECT * FROM comments
WHERE comment_id = $1;

-- name: ListCommentsByPost :many
SELECT * FROM comments
WHERE post_id = $1
ORDER BY comment_id;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE comment_id = $1;