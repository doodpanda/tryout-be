-- name: InsertUser :exec
INSERT INTO "user" (first_name, last_name, email, password)
VALUES ($1, $2, $3, $4);

-- name: LoginUser :one
SELECT id, password FROM "user"
WHERE email = $1;