-- name: CreateFollow :exec
INSERT INTO follows (
  follower_id,
  following_id
) VALUES (
  $1, $2
)
ON CONFLICT DO NOTHING;

-- name: DeleteFollow :exec
DELETE FROM follows
WHERE follower_id = $1 AND following_id = $2;

-- name: ListFollowers :many
SELECT
  u.user_id,
  u.username,
  u.full_name,
  u.profile_picture,
  u.bio,
  u.email,
  u.password_changed_at,
  u.created_at
FROM follows f
JOIN users u ON f.follower_id = u.user_id
WHERE f.following_id = $1;

-- name: ListFollowing :many
SELECT
  u.user_id,
  u.username,
  u.full_name,
  u.profile_picture,
  u.bio,
  u.email,
  u.password_changed_at,
  u.created_at
FROM follows f
JOIN users u ON f.following_id = u.user_id
WHERE f.follower_id = $1;