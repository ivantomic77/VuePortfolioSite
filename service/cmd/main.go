package main

import (
	"fmt"
	"net/http"

	"github.com/ivantomic77/VuePortfolioSite/internal"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: internal.CreateHandlers(),
	}

	fmt.Println("Starting server on :8080...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server failed:", err)
	}
}
