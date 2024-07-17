CREATE TABLE tasks (
    task_id TEXT PRIMARY KEY,
    task_description TEXT
);
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO backend;
