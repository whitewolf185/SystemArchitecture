-- +goose Up
-- +goose StatementBegin
CREATE TABLE persons (
    id uuid PRIMARY KEY NOT NULL,
    username text NOT NULL,
    user_password text NOT NULL,
    is_driver boolean NOT NULL
);
CREATE UNIQUE INDEX username_idx ON persons USING BTREE (username);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE persons;
-- +goose StatementEnd
