CREATE TABLE user_info (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    fname VARCHAR(255),
    snane VARCHAR(255),
    email citext UNIQUE NOT NULL,
    password_hash bytea NOT NULL,
    user_role VARCHAR(58),
    activated bool NOT NULL,
    version INTEGER NOT NULL DEFAULT 1
);