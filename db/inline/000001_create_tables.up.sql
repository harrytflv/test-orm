CREATE TABLE IF NOT EXISTS usersinline (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    last_name STRING,
    first_name STRING,
    nick_name STRING,
    date_of_birth STRING,
    d_id STRING,
    d_class STRING,
    d_expire_at TIMESTAMP,
    d_donor BOOL,
    p_number STRING,
    p_country STRING,
    p_signature BYTES,
    d2_id STRING,
    d2_class STRING,
    d2_expire_at TIMESTAMP,
    d2_donor BOOL,
    p2_number STRING,
    p2_country STRING,
    p2_signature BYTES
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    last_name STRING,
    first_name STRING,
    nick_name STRING,
    date_of_birth STRING,
    drivers_license JSONB,
    passport JSONB,
    drivers_license2 JSONB,
    passport2 JSONB
);

CREATE TABLE IF NOT EXISTS credit_cards (
    user_id STRING,
    card_number STRING,
    security_code STRING
)
