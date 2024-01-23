-- name: GetPost :one
SELECT * FROM posts
WHERE post_id = ? LIMIT 1;

-- name: ListPost :many
SELECT * FROM posts
ORDER BY id DESC;

-- name: CategoryPost :many
SELECT * FROM posts
WHERE category = ?;

-- name: CreatePost :execresult
INSERT INTO posts (
  post_id, title, content, category
) VALUES (
  ?, ?, ?, ? 
);

-- name: UpdatePost :execresult
UPDATE posts 
SET 
title = ?,
content = ?,
category = ?
WHERE post_id = ?; 

-- name: DeletePost :exec
DELETE FROM posts
WHERE post_id = ?;
