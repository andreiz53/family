package server

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

	database "family/database/handlers"
	"family/types"
	"family/utils"
	"family/web/components/forms"
)

func (s Server) handleRegister(w http.ResponseWriter, r *http.Request) {
	values, errors := parseAndValidateForm[types.RegisterUserValues](r)
	if *errors != (types.RegisterUserErrors{}) {
		RenderComponent(w, r, forms.RegisterForm(values, *errors))
		return
	}
	encPw := utils.EncryptPassword(values.Password)
	authUser := types.AuthenticatedUser{
		Email:    values.Email,
		LoggedIn: true,
	}
	params := types.RegisterUserValues{
		Email:    values.Email,
		Password: encPw,
	}
	_, err := s.store.CreateUser(r.Context(), types.RegisterUserToDatabaseCreateUser(params))
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			errors.Other = "User is already registered with the provided email."
			// TODO: user should have the possibility to recover his account
			RenderComponent(w, r, forms.RegisterForm(values, *errors))
			return
		}
		logError("could not register user", err)
		Redirect(w, r, "/")
		return
	}

	s.loginUser(w, r, authUser)
	Redirect(w, r, "/")
}

func (s Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	values, errors := parseAndValidateForm[types.LoginUserValues](r)
	if *errors != (types.LoginUserErrors{}) {
		RenderComponent(w, r, forms.LoginForm(values, *errors))
		return
	}

	dbUser, err := s.store.GetUserByEmail(r.Context(), values.Email)
	if err != nil {
		errors.Other = "Invalid credentials"
		RenderComponent(w, r, forms.LoginForm(values, *errors))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(values.Password))
	if err != nil {
		errors.Other = "Invalid credentials"
		RenderComponent(w, r, forms.LoginForm(values, *errors))
		return
	}
	authUser := types.AuthenticatedUser{
		Email:    values.Email,
		LoggedIn: true,
	}
	s.loginUser(w, r, authUser)
	Redirect(w, r, "/")
}

func (s Server) handleLogout(w http.ResponseWriter, r *http.Request) {
	s.logoutUser(w, r)
	Redirect(w, r, "/login")
}

func (s Server) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	u := s.getAuthenticatedUser(r)
	err := s.store.DeleteUserByEmail(r.Context(), u.Email)
	if err != nil {
		logError(fmt.Sprintf("could not delete user with email %s", u.Email), err)
		Redirect(w, r, "/")
		return
	}
	s.logoutUser(w, r)
	Redirect(w, r, "/")
}

func (s Server) handleUpdateUserEmail(w http.ResponseWriter, r *http.Request) {
	values, errors := parseAndValidateForm[types.UpdateUserEmailValues](r)
	if *errors != (types.UpdateUserEmailErrors{}) {
		RenderComponent(w, r, forms.UpdateUserEmailForm(values, *errors))
		return
	}
	authUser := s.getAuthenticatedUser(r)
	u, err := s.store.GetUserByEmail(r.Context(), authUser.Email)
	if err != nil {
		s.logoutUser(w, r)
		Redirect(w, r, "/login")
		return
	}
	err = s.store.UpdateUserEmail(r.Context(), database.UpdateUserEmailParams{
		Email: values.Email,
		ID:    u.ID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			errors.Other = "Email is already in use."
			RenderComponent(w, r, forms.UpdateUserEmailForm(values, *errors))
			return
		}
		logError("could not update user email", err)
		Redirect(w, r, "/account/settings")
		return
	}
	s.updateAuthenticatedUser(w, r, types.AuthenticatedUser{
		Email:    values.Email,
		LoggedIn: true,
	})
	Redirect(w, r, "/account/settings")
}

func (s Server) handleUpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	values, errors := parseAndValidateForm[types.UpdateUserPasswordValues](r)
	if *errors != (types.UpdateUserPasswordErrors{}) {
		RenderComponent(w, r, forms.UpdateUserPasswordForm(values, *errors))
		return
	}
	authUser := s.getAuthenticatedUser(r)
	u, err := s.store.GetUserByEmail(r.Context(), authUser.Email)
	if err != nil {
		s.logoutUser(w, r)
		Redirect(w, r, "/login")
		return
	}
	encPw := utils.EncryptPassword(values.Password)
	err = s.store.UpdateUserPassword(r.Context(), database.UpdateUserPasswordParams{
		Password: encPw,
		ID:       u.ID,
	})
	if err != nil {
		errors.Other = "Could not update password"
		RenderComponent(w, r, forms.UpdateUserPasswordForm(types.UpdateUserPasswordValues{}, *errors))
		logError("could not update user password", err)
		return
	}
	//TODO: add notification
	Redirect(w, r, "/account/settings")
}
