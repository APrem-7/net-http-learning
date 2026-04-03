package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Hello From the server")))
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
