package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write(([]byte("Hello From the index page")))
			return
		case "/users":
			w.Write(([]byte("Hello From the users page")))
			return
		case "posts":
			w.Write(([]byte("Hello From the posts page")))
			return
		}
	default:
		w.Write(([]byte("Hello From the 404 page")))
	}
}
func main() {
	api := &api{
		addr: ":8080",
	}
	s := &http.Server{
		Addr: api.addr,
	}
	if err := http.ListenAndServe(s.Addr, api); err != nil {
		log.Fatal("error")
	}

}
