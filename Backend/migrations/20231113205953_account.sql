-- +goose Up
-- +goose StatementBegin
CREATE TABLE Account(
    id serial PRIMARY KEY,
    user VARCHAR(50) UNIQUE not null,
    hash varchar(200) not null,
    registrated_at TIMESTAMP with time zone DEFAULT now(),
    is_admin BOOLEAN DEFAULT false
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Account;
-- +goose StatementEnd
