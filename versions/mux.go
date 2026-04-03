package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Hello From the index page")))
}
func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Hello From the users page")))
}
func (a *api) getPostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Hello From the posts page")))
}

func main() {
	api := &api{
		addr: ":8080",
	}
	mux := http.NewServeMux()
	s := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("GET /posts", api.getPostsHandler)
	if err := http.ListenAndServe(s.Addr, mux); err != nil {
		log.Fatal("error")
	}

}
