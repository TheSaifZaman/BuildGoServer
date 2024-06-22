package services

import (
	"fmt"
	"log"
	"net/http"
)

// HandledPrint is a helper function to write messages to the response writer
func HandledPrint(w http.ResponseWriter, m string) {
	_, err := fmt.Fprintf(w, m)
	if err != nil {
		log.Println("Error writing to response:", err)
	}
}
