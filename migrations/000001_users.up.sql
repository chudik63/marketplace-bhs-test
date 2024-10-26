CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    username VARCHAR(100) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL
);