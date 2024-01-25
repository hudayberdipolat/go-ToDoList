CREATE TYPE listStatus AS ENUM ('active', 'passive');
CREATE TABLE IF NOT EXISTS lists(
    id SERIAL PRIMARY KEY,
    list_name VARCHAR(255) NOT NULL,
    list_status listStatus ,
    user_id INT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);
