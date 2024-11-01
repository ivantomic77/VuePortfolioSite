package internal

import (
	"net/http"
)

func CreateHandlers() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/api/hello", getTest)
	return router
}
