-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email,hashed_pswd)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: Reset :exec
delete from users;


-- name: GetHashedPswd :one
select * from users where email=$1;


