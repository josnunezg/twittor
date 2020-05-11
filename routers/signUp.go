package routers

import (
	"encoding/json"
	"net/http"

	"github.com/josnunezg/twittor/bd"
	"github.com/josnunezg/twittor/models"
)

// SignUp is to create an user in the BD
func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "The data have errors: "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "The email is required", 400)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "The min of password large is 6", 400)
		return
	}

	_, foundUser, _ := bd.CheckingExistingUser(user.Email)

	if foundUser {
		http.Error(w, "The email has been used", 400)
		return
	}

	_, status, err := bd.InsertSignUp(user)

	if err != nil {
		http.Error(w, "An error ocurred on created user "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "It could not to create user", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
