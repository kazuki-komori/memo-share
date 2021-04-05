
-- +migrate Up
CREATE TABLE subjects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    teacher_id SERIAL REFERENCES teachers(id) ON DELETE SET NULL ON UPDATE CASCADE,
    semester VARCHAR(10) NOT NULL ,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
-- +migrate Down
DROP TABLE IF EXISTS subjects;