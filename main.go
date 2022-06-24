package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// content holds our static web server content.
//go:embed web/*
var content embed.FS

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	r := mux.NewRouter()

	webContent, err := fs.Sub(content, "web")
	if err != nil {
		panic("No web content found")
	}

	fs := http.FileServer(http.FS(webContent))
	r.Handle("/", fs)
	r.Handle("/index.html", fs)
	r.Handle("/favicon.ico", fs)
	r.Handle("/main.css", fs)
	r.Handle("/api/whoami", fs)
	r.Handle("/api/snapshot/new", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Tilt Cloud is deprecated", http.StatusGone)
	}))
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})

	return &Router{
		Router: r,
	}
}

func main() {
	http.Handle("/", NewRouter())

	log.Println("Serving on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Server exited with: %v", err)
	}
}
