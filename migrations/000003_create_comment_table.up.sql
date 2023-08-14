CREATE TABLE IF NOT EXISTS "comments" (
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL ,
    interview_id uuid NOT NULL,
    comment text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_interview FOREIGN KEY(interview_id) REFERENCES interviews(id)
);

CREATE INDEX comments_order_by_interview_id_and_created_at ON comments
(
    interview_id,
    created_at DESC
);