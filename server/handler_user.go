package server

import (
	"family/types"
	"family/utils"
	"family/web/components/forms"
	"net/http"
)

func (s Server) handleRegister(w http.ResponseWriter, r *http.Request) {
	values, errors := parseAndValidateForm[types.RegisterUserValues](r)
	if *errors != (types.RegisterUserErrors{}) {
		RenderComponent(w, r, forms.RegisterForm(values, *errors))
		return
	}
	// encrypt password
	encPw := utils.EncryptPassword(values.Password)
	// insert into db
	// check for error and duplicate key
	_ = encPw

	user := types.AuthenticatedUser{
		Email:    values.Email,
		LoggedIn: true,
	}
	s.loginUser(w, r, user)
	Redirect(w, r, "/")
}

func (s Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	values, errors := parseAndValidateForm[types.LoginUserValues](r)
	if *errors != (types.LoginUserErrors{}) {
		RenderComponent(w, r, forms.LoginForm(values, *errors))
		return
	}

	// get the user from db
	// compare passwords

	user := types.AuthenticatedUser{
		Email:    values.Email,
		LoggedIn: true,
	}
	s.loginUser(w, r, user)
	Redirect(w, r, "/")
}

func (s Server) handleLogout(w http.ResponseWriter, r *http.Request) {
	s.logoutUser(w, r)
	Redirect(w, r, "/login")
}
