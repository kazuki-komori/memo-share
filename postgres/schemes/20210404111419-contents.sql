
-- +migrate Up
CREATE TABLE contents (
    id VARCHAR(50) PRIMARY KEY,
    content_title VARCHAR(100) NOT NULL,
    description VARCHAR(200) NOT NULL,
    user_id VARCHAR(100) REFERENCES users(id) ON DELETE SET NULL ON UPDATE CASCADE,
    subject_id VARCHAR(50) NOT NULL ,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
-- +migrate Down
DROP TABLE IF EXISTS contents;