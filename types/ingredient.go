package types

import (
	"github.com/jackc/pgx/v5/pgtype"

	database "family/database/handlers"
)

type Ingredient struct {
	ID        pgtype.UUID
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	Name      string
}

func DBIngredientToIngredient(dbIngredient database.Ingredient) Ingredient {
	return Ingredient{
		ID:        dbIngredient.ID,
		CreatedAt: dbIngredient.CreatedAt,
		UpdatedAt: dbIngredient.UpdatedAt,
		Name:      dbIngredient.Name,
	}
}

func DBIngredientsToIngredients(dbItems []database.Ingredient) []Ingredient {
	ingredients := []Ingredient{}
	for _, item := range dbItems {
		ingredients = append(ingredients, DBIngredientToIngredient(item))
	}
	return ingredients
}
