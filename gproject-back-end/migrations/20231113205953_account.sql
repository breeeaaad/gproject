-- +goose Up
-- +goose StatementBegin
CREATE TABLE Account(
    id serial PRIMARY KEY,
    username VARCHAR(50) UNIQUE not null,
    hash varchar(200) not null,
    registrated_at TIMESTAMP with time zone DEFAULT now(),
    is_admin BOOLEAN DEFAULT false,
    tfa CHAR(26) UNIQUE
);
CREATE INDEX username on Account using hash (username);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Account;
-- +goose StatementEnd
