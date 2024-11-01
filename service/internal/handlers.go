package internal

import (
	"fmt"
	"net/http"
)

func getTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
