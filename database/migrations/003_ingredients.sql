-- +goose Up
CREATE TABLE ingredients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
    name TEXT NOT NULL UNIQUE
);



-- +goose Down
DROP TABLE ingredients;
