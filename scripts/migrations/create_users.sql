CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS users (
    id  UUID DEFAULT gen_random_uuid(),
    login TEXT not null,
    email TEXT not null,
    created_at timestamp not null,
    updated_at timestamp not null
);