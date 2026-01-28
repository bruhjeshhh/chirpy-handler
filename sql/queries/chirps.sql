-- name: PostChirp :one
INSERT INTO chirps(id, created_at, updated_at, body, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: GetChirps :many
select * from chirps;

-- name: GetChirpsbyID :one
select * from chirps where id=$1;



-- name: DeleteChirp :exec
delete from chirps where id=$1 and user_id=$2;

-- name: GetChirpsByAuthor :many
select created_at,body from chirps where user_id=$1;