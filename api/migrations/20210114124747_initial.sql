-- +goose Up
-- SQL in this section is executed when the migration is applied.
BEGIN;
CREATE TABLE IF NOT EXISTS projects (
    id UUID DEFAULT uuid_generate_v4(),
    user_id varchar NOT NULL,

    pname varchar NOT NULL UNIQUE,
    pdescription varchar,


    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id)
    );
COMMIT;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

BEGIN;
DROP TABLE projects;
COMMIT;