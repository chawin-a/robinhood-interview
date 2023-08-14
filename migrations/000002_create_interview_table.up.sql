CREATE TYPE interview_status AS ENUM ('TO_DO', 'IN_PROGRESS', 'DONE');

CREATE TABLE IF NOT EXISTS "interviews" (
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL,
    description text NOT NULL,
    status interview_status NOT NULL DEFAULT 'TO_DO'::interview_status,
    is_archived boolean NOT NULL DEFAULT false,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) 
);

CREATE INDEX interview_order_by_created_at ON interviews (created_at);

CREATE TRIGGER update_interview_updated_at
    AFTER UPDATE
    ON
        interviews
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at();