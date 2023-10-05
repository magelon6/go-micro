package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) Auth(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email 		string 	`json:"email"`
		Password 	string 	`json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := app.Models.User.PasswordMatches(user.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("Invalid password"), http.StatusUnauthorized)
		return
	}

	payload := jsonResponse {
		Error: false,
		Message: fmt.Sprintf("Logged as %s", user.Email),
		Data: user,
	}

	app.writeJSON(w, http.StatusOK, payload)

}