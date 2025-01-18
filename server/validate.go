package server

import (
	"family/utils"
	"log/slog"
	"net/http"
)

type ValidatableForm[E FormError] interface {
	Validate() E
}

type FormError interface {
	SetOtherError(message string)
}

func parseAndValidateForm[V ValidatableForm[E], E FormError](r *http.Request) (V, E) {
	var values V
	var errors E
	err := r.ParseForm()
	if err != nil {
		errors.SetOtherError("Something went wrong, please try again.")
		return values, errors
	}

	var formData = make(map[string]interface{})
	for k, v := range r.Form {
		if len(v) > 0 {
			formData[k] = v[0]
		}
	}

	err = utils.FillStructFromMap(&values, formData)
	if err != nil {
		logError("could transform map to struct in parsing the form", err)
		errors.SetOtherError("Something went wrong, please try again.")
		return values, errors
	}

	errors = values.Validate()
	return values, errors
}

func logError(message string, err error) {
	slog.Error(message, "error", err.Error())
}
