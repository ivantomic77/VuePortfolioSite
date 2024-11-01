package internal

import (
	"net/http"
)

func CreateHandlers() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/api/posts/{slug}", getPost)
	router.HandleFunc("/api/previews", getPreviews)
	return router
}
