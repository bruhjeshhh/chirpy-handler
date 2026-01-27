-- +goose Up
CREATE TABLE users(
	id uuid primary key,
	created_at timestamp not null,
	updated_at timestamp not null,
	email text unique not null,
	hashed_pswd text not null,
	is_chirpy_red boolean default false
);

-- +goose Down
DROP TABLE users;


 