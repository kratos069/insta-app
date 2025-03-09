-- name: CreateLike :one
INSERT INTO likes (
  user_id, post_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: CountLikesByPost :many
SELECT COUNT(*) AS count
FROM likes
WHERE post_id = $1;

-- name: DeleteLike :exec
DELETE FROM likes
WHERE post_id = $1 AND user_id = $2;