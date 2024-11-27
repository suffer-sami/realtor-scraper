-- +goose Up
CREATE TABLE agents (
    id TEXT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,

    first_name TEXT,
    last_name TEXT,
    nick_name TEXT,
    person_name TEXT,
    title TEXT,
    slogan TEXT,

    email TEXT,

    agent_rating INTEGER,

    description TEXT,
    recommendations_count INTEGER,
    review_count INTEGER,

    last_updated TIMESTAMP,
    first_month INTEGER,
    first_year INTEGER,

    video TEXT,
    web_url TEXT,

    href TEXT
);

-- +goose Down
DROP TABLE agents;