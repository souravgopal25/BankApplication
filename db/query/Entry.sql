-- name: CreateEntry :one
INSERT INTO entries(
        account_id, amount
) values (
          $1 , $2
          )
returning *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: GetEntriesOfAccount :many
SELECT * FROM entries
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateEntry :one
UPDATE entries
SET account_id = $1,
    amount=$2
WHERE id=$3
RETURNING *;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;



