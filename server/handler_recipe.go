package server

import (
	"net/http"
	"strings"

	"family/types"
	"family/web/components/forms"
)

func (s Server) handleCreateRecipe(w http.ResponseWriter, r *http.Request) {}

func (s Server) handleCreateRecipeItem(w http.ResponseWriter, r *http.Request) {
	values, errors := parseAndValidateForm[types.CreateRecipeItemValues](r)
	if *errors != (types.CreateRecipeItemErrors{}) {
		RenderComponent(w, r, forms.CreateRecipeItemForm(values, *errors))
		return
	}
	err := s.store.CreateRecipeItem(r.Context(), values.Name)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			errors.Other = "Ingredient already exists."
			RenderComponent(w, r, forms.CreateRecipeItemForm(values, *errors))
			return
		}
		logError("could not create recipe item", err)
		Redirect(w, r, "/ingredients/new")
		return
	}
	RenderComponent(w, r, forms.CreateRecipeItemForm(types.CreateRecipeItemValues{}, types.CreateRecipeItemErrors{}))
}
