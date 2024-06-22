package main

import (
	"BuildGoServer/internal/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Serving static files from the ./static directory
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Handlers for different endpoints
	http.HandleFunc("/form", handlers.FormHandler)
	http.HandleFunc("/hello", handlers.HelloHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
