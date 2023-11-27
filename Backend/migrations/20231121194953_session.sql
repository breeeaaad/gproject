-- +goose Up
-- +goose StatementBegin
create table Session(
    id serial PRIMARY key,
    user_id int REFERENCES account(id) on delete cascade,
    refresh uuid default gen_random_uuid(),
    expiresIn TIMESTAMP with time zone DEFAULT now() + interval '7 days'
);
CREATE index refresh on Session using hash (refresh);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if EXISTS Session;
-- +goose StatementEnd
