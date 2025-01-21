-- +goose Up
CREATE TABLE recipe_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
    deleted_at TIMESTAMP, 
    name TEXT NOT NULL UNIQUE 
);


-- +goose Down
DROP TABLE recipe_items;
