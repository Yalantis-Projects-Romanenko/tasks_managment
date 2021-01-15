-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS projects (
    id UUID DEFAULT uuid_generate_v4(),
    user_id varchar NOT NULL,

    pname varchar NOT NULL UNIQUE,
    pdescription varchar,


    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS columns (
    id UUID DEFAULT uuid_generate_v4(),
    project_id varchar NOT NULL,
    cname varchar NOT NULL ,
    index integer ,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS tasks (
    id UUID DEFAULT uuid_generate_v4(),
    title varchar NOT NULL,
    description varchar NOT NULL,
    priority integer,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS comments (
    id UUID DEFAULT uuid_generate_v4(),
    username varchar NOT NULL,
    content varchar NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id)
    );



-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE projects;
DROP TABLE columns;
DROP TABLE tasks;
DROP TABLE comments;
