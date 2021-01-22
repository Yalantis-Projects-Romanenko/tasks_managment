-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE tasks
    ADD COLUMN project_Id UUID;

UPDATE tasks
SET project_Id = lists.project_Id
FROM lists
WHERE lists.id = tasks.column_id;

ALTER TABLE lists
    ALTER COLUMN project_id SET NOT NULL;

ALTER TABLE comments
    ADD COLUMN project_Id UUID;

UPDATE comments
SET project_Id = tasks.project_Id
FROM tasks
WHERE tasks.id = comments.task_id;

ALTER TABLE comments
    ALTER COLUMN project_id SET NOT NULL;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE tasks
    DROP COLUMN project_Id;
ALTER TABLE comments
    DROP COLUMN project_Id;
