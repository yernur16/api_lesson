-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY NOT NULL,
    data VARCHAR) ;


-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
