package types

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgtype"

	database "family/database/handlers"
	"family/utils"
	"family/validate"
)

type Recipe struct {
	ID             pgtype.UUID
	Name           string
	CookingProcess string
	Description    string
	Ingredients    []RecipeItem
}

type RecipeItem struct {
	ID        pgtype.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type CreateRecipeValues struct {
	Name           string
	CookingProcess string
	Description    string
	Ingredients    string
}

type CreateRecipeErrors struct {
	Name           string
	CookingProcess string
	Description    string
	Other          string
}

type CreateRecipeItemValues struct {
	Name string `validate:"required"`
}

type CreateRecipeItemErrors struct {
	Name  string
	Other string
}

func (values CreateRecipeItemValues) Validate() *CreateRecipeItemErrors {
	errors := &CreateRecipeItemErrors{}

	v := validate.Validator
	err := v.Struct(values)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "Name":
					errors.Name = "Field is required."
				}
			}
		}
	}

	return errors
}

func (errors *CreateRecipeItemErrors) SetOtherError(message string) {
	errors.Other = message
}

func DBRecipeItemToRecipeItem(dbItem database.RecipeItem) RecipeItem {
	return RecipeItem{
		ID:        dbItem.ID,
		CreatedAt: dbItem.CreatedAt.Time,
		UpdatedAt: dbItem.UpdatedAt.Time,
		DeletedAt: utils.NullTime(dbItem.DeletedAt),
		Name:      dbItem.Name,
	}
}

func DBRecipeItemsToRecipeItems(dbItems []database.RecipeItem) []RecipeItem {
	recipeItems := []RecipeItem{}
	for _, item := range dbItems {
		recipeItems = append(recipeItems, DBRecipeItemToRecipeItem(item))
	}

	return recipeItems
}
