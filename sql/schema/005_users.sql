-- +goose up
alter table users
add column is_chirpy_red boolean default false;

-- +goose down
alter table users
drop column is_chirpy_red;