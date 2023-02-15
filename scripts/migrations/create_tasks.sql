CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS tasks (
    id  UUID DEFAULT gen_random_uuid(),
    user_id UUID,
    name TEXT not null ,
    created_at timestamp not null,
    updated_at timestamp not null
);