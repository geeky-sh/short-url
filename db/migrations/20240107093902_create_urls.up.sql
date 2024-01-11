CREATE TABLE IF NOT EXISTS urls (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INT NOT NULL REFERENCES users,
    code VARCHAR(32) NOT NULL UNIQUE,
    url VARCHAR(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp
);
