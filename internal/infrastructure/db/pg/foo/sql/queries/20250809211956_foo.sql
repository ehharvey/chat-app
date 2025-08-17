-- name: InsertOneFoo :one
INSERT INTO foo (name) VALUES
($1) RETURNING *;

-- name: CheckExistsFooByName :one
SELECT EXISTS(
    SELECT 1 FROM foo WHERE name = $1 LIMIT 1
) AS exists;