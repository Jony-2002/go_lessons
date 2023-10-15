package controllers

import (
	"encoding/json"
	"khorogtech/types"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var newUser types.User

	json.NewDecoder(r.Body).Decode(&newUser)

	// check if user already exist with the same email
	// check the empty fields
	// check the password - more than 6, first letter uppercase, with special character
	// check Email character with @gmail.com | @email.com
}