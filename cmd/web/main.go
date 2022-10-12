package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	var (
		port = flags.Int("port", 80, "port to listen on")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}
	addr := fmt.Sprintf(":%d", *port)

	srv := newServer()
	fmt.Printf("Starting front-end, listening on :%d\n", *port)
	return http.ListenAndServe(addr, srv)
}

type server struct {
	mux *http.ServeMux
}

func newServer() *server {
	mux := http.NewServeMux()

	srv := &server{
		mux: mux,
	}
	srv.routes()
	return srv
}

func (s *server) routes() {
	s.mux.Handle("/signin", s.handleSignIn())
	s.mux.Handle("/signup", s.handleSignUp())
	s.mux.Handle("/", s.handleIndex())
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, "index")
	}
}

func (s *server) handleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, "signin")
	}
}

func (s *server) handleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, "signup")
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
