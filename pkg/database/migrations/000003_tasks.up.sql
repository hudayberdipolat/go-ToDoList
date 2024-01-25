CREATE TYPE taskStatus AS ENUM ('active', 'passive');
CREATE TABLE IF NOT EXISTS tasks(
    id SERIAL PRIMARY KEY,
    task_name VARCHAR(255) NOT NULL,
    task_status taskStatus,
    list_id INT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_list FOREIGN KEY(list_id) REFERENCES lists(id) ON DELETE CASCADE
);