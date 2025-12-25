package main

import (
	"log"
	"net/http"
)

type server struct {
	addr string
}

func (s *server) getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Luare User"))
}

func (s *server) createUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creare User"))
}

func main() {
	s := &server{addr: ":8080"}
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}
	mux.HandleFunc("GET /users", s.getUserHandler)
	mux.HandleFunc("POST /users", s.createUserHandler)
	log.Fatal(http.ListenAndServe(srv.Addr, srv.Handler))
}
