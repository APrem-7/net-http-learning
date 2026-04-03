package main

import (
	"log"
	"net/http"
)

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
	mux.HandleFunc("POST /users", api.getPostsHandler)
	if err := http.ListenAndServe(s.Addr, mux); err != nil {
		log.Fatal(err)

	}

}
