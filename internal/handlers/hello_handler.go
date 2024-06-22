package handlers

import (
	"BuildGoServer/internal/utils"
	"net/http"
)

// HelloHandler handles the GET request for the /hello endpoint
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	utils.HandledPrint(w, "Hello World!")
}
