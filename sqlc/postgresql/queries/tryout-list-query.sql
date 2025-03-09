-- name: GetTryoutListFiltered :many
SELECT * FROM tryout
WHERE (is_published = true OR creator_id = $1)
AND difficulty = $2
AND category = $3
AND title ILIKE '%' || $4 || '%'
ORDER BY created_at DESC;

-- name: GetTryoutList :many
SELECT * FROM tryout
WHERE (is_published = true OR creator_id = $1)
ORDER BY created_at DESC;

-- name: GetTryoutById :one
SELECT * FROM tryout
WHERE id = $1;