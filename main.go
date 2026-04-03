package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Hello From the index page")))
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

	mux.HandleFunc("GET /users", api.getUsersHandler) //still not implenented
	mux.HandleFunc("GET /posts", api.getPostsHandler) //still not implemetned
	if err := http.ListenAndServe(s.Addr, api); err != nil {
		log.Fatal("error")
	}

}
