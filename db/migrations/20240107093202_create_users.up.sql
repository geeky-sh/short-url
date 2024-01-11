CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(32) NOT NULL UNIQUE,
    encrypted_password VARCHAR(255) NOT NULL,
    created_at timestamp NOT NULL,
    last_logged_in_at timestamp NULL
);
