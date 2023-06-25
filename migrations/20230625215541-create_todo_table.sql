
-- +migrate Up
CREATE TABLE IF NOT EXISTS todos (
    id uuid PRIMARY KEY,
    title varchar(100),
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
-- +migrate Down
DROP TABLE IF EXISTS todos;
