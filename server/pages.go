package server

import (
	"net/http"

	"family/types"
	"family/web/pages"
	accountPages "family/web/pages/account"
	ingredientsPages "family/web/pages/ingredients"
	recipePages "family/web/pages/recipe"
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
	u := s.getAuthenticatedUser(r)

	dbUser, err := s.store.GetUserByEmail(r.Context(), u.Email)
	if err != nil {
		s.logoutUser(w, r)
		Redirect(w, r, "/login")
	}
	RenderComponent(w, r, accountPages.AccountSettings(types.DBUserToUser(dbUser)))
}

func (s Server) renderRecipeCreate(w http.ResponseWriter, r *http.Request) {
	RenderComponent(w, r, recipePages.CreateRecipe())
}

func (s Server) renderIngredientsIndex(w http.ResponseWriter, r *http.Request) {
	recipeItems, err := s.store.GetRecipeItems(r.Context())
	if err != nil {
		logError("could not get ingredients from store", err)
		Redirect(w, r, "/ingredients/new")
		return
	}
	RenderComponent(w, r, ingredientsPages.IngredientsIndex(types.DBRecipeItemsToRecipeItems(recipeItems)))
}

func (s Server) renderIngredientCreate(w http.ResponseWriter, r *http.Request) {
	RenderComponent(w, r, ingredientsPages.IngredientCreate())
}
