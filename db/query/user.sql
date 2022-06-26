-- name: CreateUser :one
INSERT INTO users(USERNAME,
                  HASHED_PASSWORD,
                  FULL_NAME,
                  EMAIL)
VALUES ($1, $2, $3, $4)
RETURNING *;


-- name: GetUser :one
SELECT *
FROM users
WHERE username = $1
LIMIT 1;