-- name: PostChirp :one
insert into chirps (id, created_at, updated_at, body, user_id)
values ($1, $2, $3, $4, $5)
returning *;

-- name: AllChirps :many
select * from chirps
order by created_at;

-- name: GetChirpByChirpId :one
select * from chirps
where id = $1;