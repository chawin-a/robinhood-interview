CREATE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS users (
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    username text NOT NULL UNIQUE,
    name text NOT NULL,
    email text NOT NULL,
    created_at timestamptz DEFAULT NOW()
);