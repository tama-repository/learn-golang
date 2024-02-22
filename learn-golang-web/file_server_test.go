package learn_golang_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	dir := http.Dir("./resources")

	fileServer := http.FileServer(dir)

	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

//go:embed resources
var resources embed.FS

func TestFileServerEmbed(t *testing.T) {

	dir, err := fs.Sub(resources, "resources")

	if err != nil {
		panic(err)
	}

	fileServer := http.FileServer(http.FS(dir))

	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err2 := server.ListenAndServe()

	if err2 != nil {
		panic(err2)
	}
}
