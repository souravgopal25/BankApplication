
-- name: CreateAccount :one
INSERT INTO account (
    owner, balnce, currency
) VALUES (
             $1, $2, $3
         )
RETURNING *;