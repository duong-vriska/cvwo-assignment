-- name: GetPost :one
SELECT * FROM posts
WHERE id = ? LIMIT 1;

-- name: ListPost :many
SELECT * FROM posts
ORDER BY id DESC;

-- name: CategoryPost :many
SELECT * FROM posts
WHERE category = ?;

-- name: CreatePost :execresult
INSERT INTO posts (
  id, post_id, title, content, category
) VALUES (
  ?, ?, ?, ?, ? 
);

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = ?;

