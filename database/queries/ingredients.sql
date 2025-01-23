-- name: CreateIngredient :exec
INSERT INTO ingredients (name)
VALUES ($1);


-- name: GetIngredientByID :one
SELECT * FROM ingredients
WHERE id = $1;

-- name: GetIngredients :many
SELECT * FROM ingredients;

-- name: UpdateIngredient :exec
UPDATE ingredients SET
name = $1,
updated_at = NOW()
WHERE id = $2;

-- name: DeleteIngredient :exec
DELETE FROM ingredients
WHERE id = $1;

-- name: DeleteAllIngredients :exec
DELETE FROM ingredients;
