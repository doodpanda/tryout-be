-- name: GetTryoutListFiltered :many
SELECT * FROM tryout
WHERE (is_published = true OR (creator_id IS NOT NULL AND creator_id = NULLIF($1, '')::UUID))
AND (difficulty = $2 OR $2 IS NULL)
AND (category = $3 OR $3 IS NULL)
AND (title ILIKE '%' || $4 || '%' OR $4 IS NULL)
ORDER BY created_at DESC;

-- name: GetTryoutList :many
SELECT * FROM tryout
WHERE (is_published = true OR (creator_id IS NOT NULL AND creator_id = NULLIF($1, '')::UUID))
ORDER BY created_at DESC;

-- name: GetTryoutCreator :one
SELECT creator_id FROM tryout
WHERE id = $1;

-- name: GetTryoutById :one
SELECT * FROM tryout
WHERE id = $1;

-- name: InsertTryout :exec
INSERT INTO tryout (creator_id, title, description, long_description, difficulty, duration, topics, category, is_published, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW());

-- name: UpdateTryout :exec
UPDATE tryout
SET title = $2,
    description = $3,
    long_description = $4,
    difficulty = $5,
    category = $6,
    is_published = $7,
    duration = $8,
    topics = $9
WHERE id = $1;

-- name: DeleteTryout :exec
DELETE FROM tryout
WHERE id = $1;
