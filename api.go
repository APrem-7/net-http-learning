package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type api struct {
	addr string
}

var users = []Users{}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Hello From the index page")))

}
func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//now we are getting the users slice so

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (a *api) getPostsHandler(w http.ResponseWriter, r *http.Request) {
	var payload User
	w.Header().Set("Content-Type", "application/json")
	//now we are posting into the users slice so

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u := User{
		first_name: payload.first_name,
		last_name:  payload.last_name,
	}

	users = append(users, u)

	w.WriteHeader(http.StatusOK)
}

func InsertUser(u User) error {
	if u.first_name == "" {
		return errors.New("first_name cannot be empty")
	}
	if u.last_name == "" {
		return errors.New("last_name cannot be empty")
	}
	return nil
}
