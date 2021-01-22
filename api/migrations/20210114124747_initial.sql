-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS projects
(
    id           UUID               DEFAULT uuid_generate_v4(),
    user_id      varchar   NOT NULL,
    pname        varchar   NOT NULL,
    pdescription varchar,


    created_at   TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS lists
(
    id         UUID               DEFAULT uuid_generate_v4(),
    project_id UUID      NOT NULL,
    cname      varchar   NOT NULL,
    index      integer,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id),
    CONSTRAINT fk_project
        FOREIGN KEY (project_id)
            REFERENCES projects (id)
            ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS tasks
(
    id          UUID               DEFAULT uuid_generate_v4(),
    column_id   UUID      NOT NULL,
    title       varchar   NOT NULL,
    description varchar   NOT NULL,
    priority    integer,

    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id),
    CONSTRAINT fk_column
        FOREIGN KEY (column_id)
            REFERENCES lists (id)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS comments
(
    id         UUID               DEFAULT uuid_generate_v4(),
    task_id    UUID      NOT NULL,
    username   varchar   NOT NULL,
    content    varchar   NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id),
    CONSTRAINT fk_task
        FOREIGN KEY (task_id)
            REFERENCES tasks (id)
            ON DELETE CASCADE
);



-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE projects;
DROP TABLE lists;
DROP TABLE tasks;
DROP TABLE comments;
