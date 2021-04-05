
-- +migrate Up
CREATE TABLE teachers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);
-- +migrate Down
DROP TABLE IF EXISTS teachers;