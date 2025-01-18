package server

import (
	"family/web/pages"
	accountPages "family/web/pages/account"
	"net/http"
)

func (s Server) renderIndex(w http.ResponseWriter, r *http.Request) {
	RenderComponent(w, r, pages.Index())
}

func (s Server) renderLogin(w http.ResponseWriter, r *http.Request) {
	RenderComponent(w, r, pages.Login())
}

func (s Server) renderRegister(w http.ResponseWriter, r *http.Request) {
	RenderComponent(w, r, pages.Register())
}

func (s Server) renderAccountSettings(w http.ResponseWriter, r *http.Request) {
	RenderComponent(w, r, accountPages.AccountSettings())
}
