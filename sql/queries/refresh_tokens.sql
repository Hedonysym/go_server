-- name: NewRefreshToken :one
insert into refresh_tokens (token, created_at, updated_at, user_id, expires_at)
values ($1, $2, $3, $4, $5)
on conflict (token) do update
set updated_at = $3, expires_at = $5
returning *;

-- name: GetUserByRefreshToken :one
SELECT rt.token, rt.expires_at, rt.revoked_at, u.id, u.email, u.created_at, u.updated_at
FROM refresh_tokens rt
JOIN users u ON rt.user_id = u.id
WHERE rt.token = $1;

-- name: RevokeRefreshToken :exec
update refresh_tokens
set revoked_at = $2, updated_at = $3
where token = $1;

-- name: GetRefreshToken :one
select * from refresh_tokens
where token = $1;