CREATE TYPE listStatus AS ENUM ('1', '0');
CREATE TABLE IF NOT EXISTS lists(
    id SERIAL PRIMARY KEY,
    list_name VARCHAR(255) NOT NULL,
    list_status listStatus DEFAULT '1',
    user_id INT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);
