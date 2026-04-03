package main

import (
	"encoding/json"
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

	users.append(u)

	w.WriteHeader(http.StatusOK)
}
