package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	database "family/database/handlers"
)

type Storage interface {
	UserStorage
	RecipeItemStorage
	IngredientStorage
}

type UserStorage interface {
	CreateUser(context.Context, database.CreateUserParams) (database.User, error)
	GetUsers(context.Context) ([]database.User, error)
	GetUserByEmail(context.Context, string) (database.User, error)
	GetUserByID(context.Context, pgtype.UUID) (database.User, error)
	DeleteUserByID(context.Context, pgtype.UUID) error
	DeleteUserByEmail(context.Context, string) error
	UpdateUserEmail(context.Context, database.UpdateUserEmailParams) error
	UpdateUserPassword(context.Context, database.UpdateUserPasswordParams) error
}

type RecipeItemStorage interface {
	CreateRecipeItem(context.Context, string) error
	DeleteRecipeItem(context.Context, pgtype.UUID) error
	GetRecipeItemByID(context.Context, pgtype.UUID) (database.RecipeItem, error)
	GetRecipeItems(context.Context) ([]database.RecipeItem, error)
	UpdateRecipeItem(context.Context, string) error
}

type IngredientStorage interface {
	CreateIngredient(context.Context, string) error
	GetIngredientByID(context.Context, pgtype.UUID) (database.Ingredient, error)
	GetIngredients(context.Context) ([]database.Ingredient, error)
	UpdateIngredient(context.Context, database.UpdateIngredientParams) error
	DeleteIngredient(context.Context, pgtype.UUID) error
	DeleteAllIngredients(context.Context) error
}

type Database struct {
	Queries *database.Queries
}

func NewDatabase(url string) Database {

	db, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatal("could not connect to database", err)
	}
	queries := database.New(db)

	return Database{
		Queries: queries,
	}
}
