
-- +migrate Up
CREATE TABLE teachers (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);
-- +migrate Down
DROP TABLE IF EXISTS teachers;