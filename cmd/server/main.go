package main

import (
	"family/cookies"
	"family/server"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cookieStore := cookies.NewCookieStore(os.Getenv("SESSION_KEY"))
	server := server.NewServer(os.Getenv("PORT"), cookieStore)
	server.Init()
}
