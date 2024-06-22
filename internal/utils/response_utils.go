package utils

import (
	"BuildGoServer/internal/services"
	"net/http"
)

// HandledPrint is a utility wrapper for the print service
func HandledPrint(w http.ResponseWriter, m string) {
	services.HandledPrint(w, m)
}
