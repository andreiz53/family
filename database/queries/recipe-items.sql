-- name: CreateRecipeItem :exec
INSERT INTO recipe_items (name)
VALUES ($1);

-- name: GetRecipeItemByID :one
SELECT * FROM recipe_items
WHERE id = $1
AND deleted_at IS NULL;

-- name: GetRecipeItems :many
SELECT * FROM recipe_items
WHERE deleted_at IS NULL;

-- name: UpdateRecipeItem :exec
UPDATE recipe_items
SET name = $1,
updated_at = NOW()
WHERE id = $1;

-- name: DeleteRecipeItem :exec
UPDATE recipe_items
SET deleted_at = NOW()
WHERE id = $1;
