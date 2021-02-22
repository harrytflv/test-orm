CREATE TABLE IF NOT EXISTS users (
    id STRING PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    last_name STRING,
    first_name STRING,
    nick_name STRING,
    date_of_birth STRING,
    drivers_license JSONB,
    passport JSONB
);

CREATE TABLE IF NOT EXISTS credit_cards (
    user_id STRING,
    card_number STRING,
    security_code STRING
)
