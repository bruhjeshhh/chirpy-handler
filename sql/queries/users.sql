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


-- name: UpdateEmail :exec
update users set email=$1,hashed_pswd=$2,updated_at=$3
where id =$4;

