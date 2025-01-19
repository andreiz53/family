package main

import (
	"family/cookies"
	"family/database"
	"family/server"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cookieStore := cookies.NewCookieStore(os.Getenv("SESSION_KEY"))
	database := database.NewDatabase(os.Getenv("DATABASE_URL"))
	server := server.NewServer(os.Getenv("PORT"), cookieStore, database.Queries)
	server.Init()
}
