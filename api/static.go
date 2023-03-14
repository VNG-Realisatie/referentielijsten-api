package api

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// serveRedoc serves the static redocly to include documentation with the api
func serveRedoc(w http.ResponseWriter, req *http.Request) {
	name := "docs/redoc.html"
	absolutePath, _ := filepath.Abs(name)

	if _, err := os.Stat(absolutePath); errors.Is(err, os.ErrNotExist) {
		absolutePath = filepath.Join("/app", name)
	}

	log.Println(absolutePath)
	http.ServeFile(w, req, absolutePath)
}

// serveOpenApi makes the openapi.yaml downloadable
func serveOpenApi(w http.ResponseWriter, req *http.Request) {
	name := "docs/openapi.yaml"
	absolutePath, _ := filepath.Abs(name)

	if _, err := os.Stat(absolutePath); errors.Is(err, os.ErrNotExist) {
		absolutePath = filepath.Join("/app", name)
	}

	openApi, err := os.ReadFile(absolutePath)
	if err != nil {
		log.Println(err)
	}

	w.Write(openApi)
}
