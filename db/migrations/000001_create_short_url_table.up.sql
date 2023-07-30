CREATE TABLE IF NOT EXISTS short_urls (
    id SERIAL PRIMARY KEY,
    code VARCHAR(32),
    url VARCHAR(255),
    created_at timestamp DEFAULT current_timestamp
);
