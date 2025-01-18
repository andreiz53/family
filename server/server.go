package server

import (
	"fmt"
	"family/cookies"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	listenAddr  string
	cookieStore cookies.CookieStore
}

func NewServer(addr string, cs cookies.CookieStore) Server {
	return Server{
		listenAddr:  addr,
		cookieStore: cs,
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

	router.Mount("/account", accountRouter)

	fmt.Printf("Server starting on port %s\n", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}
