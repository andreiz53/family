package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"family/cookies"
	"family/database"
)

type Server struct {
	listenAddr  string
	cookieStore cookies.CookieStore
	store       database.Storage
}

func NewServer(addr string, cs cookies.CookieStore, store database.Storage) Server {
	return Server{
		listenAddr:  addr,
		cookieStore: cs,
		store:       store,
	}
}

func (s Server) Init() {

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(s.WithUser)

	fs := http.FileServer(http.Dir("web/assets"))
	router.Handle("/assets/*", http.StripPrefix("/assets", fs))

	router.Get("/", s.renderIndex)
	router.Get("/login", s.renderLogin)
	router.Get("/register", s.renderRegister)
	router.Post("/login", s.handleLogin)
	router.Post("/register", s.handleRegister)

	accountRouter := chi.NewRouter()
	accountRouter.Use(s.WithAuth)
	accountRouter.Post("/logout", s.handleLogout)
	accountRouter.Get("/settings", s.renderAccountSettings)
	accountRouter.Put("/settings/email", s.handleUpdateUserEmail)
	accountRouter.Put("/settings/password", s.handleUpdateUserPassword)
	accountRouter.Delete("/", s.handleDeleteUser)

	router.Mount("/account", accountRouter)

	recipeRouter := chi.NewRouter()
	recipeRouter.Get("/", s.renderRecipesIndex)
	recipeRouter.Get("/create", s.renderRecipeCreate)

	router.Mount("/recipes", recipeRouter)

	ingredientRouter := chi.NewRouter()
	ingredientRouter.Get("/", s.renderIngredientsIndex)

	router.Mount("/ingredients", ingredientRouter)

	fmt.Printf("Server starting on port %s\n", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}
