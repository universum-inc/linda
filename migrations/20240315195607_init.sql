-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS account (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(32) UNIQUE NOT NULL,
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(128),
    date_of_birth DATE
);

CREATE TABLE IF NOT EXISTS account_contact_info (
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT REFERENCES account(id) ON DELETE CASCADE ON UPDATE CASCADE,
    contact_type VARCHAR(50),
    contact_details TEXT
);

CREATE TABLE IF NOT EXISTS photo (
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT REFERENCES account(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    link TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS video (
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT REFERENCES account(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    link TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists video;
drop table if exists photo;
drop table if exists account_contact_info;
drop table if exists account;
-- +goose StatementEnd
