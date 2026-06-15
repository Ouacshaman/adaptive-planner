-- +goose Up
CREATE TABLE tasks(
id UUID PRIMARY KEY,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL,
task TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE tasks;
