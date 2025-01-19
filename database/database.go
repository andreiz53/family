package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	database "family/database/handlers"
)

type Storage interface {
	CreateUser(context.Context, database.CreateUserParams) (database.User, error)
	GetUsers(context.Context) ([]database.User, error)
	GetUserByEmail(context.Context, string) (database.User, error)
	GetUserByID(context.Context, pgtype.UUID) (database.User, error)
	DeleteUserByID(context.Context, pgtype.UUID) error
	DeleteUserByEmail(context.Context, string) error
	UpdateUserEmail(context.Context, database.UpdateUserEmailParams) error
	UpdateUserPassword(context.Context, database.UpdateUserPasswordParams) error
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
