
-- +migrate Up
CREATE TABLE users(
                          id VARCHAR(100) PRIMARY KEY NOT NULL,
                          mail VARCHAR(100) NOT NULL,
                          user_name VARCHAR(50) NOT NULL,
                          full_name VARCHAR(50) NOT NULL,
                          created_at TIMESTAMP NOT NULL ,
                          updated_at TIMESTAMP
);
-- +migrate Down
DROP TABLE IF EXISTS users;