package handlers

import (
	"BuildGoServer/internal/utils"
	"net/http"
)

// FormHandler handles the POST request for the form submission
func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.NotFound(w, r)
		return
	}

	if err := r.ParseForm(); err != nil {
		utils.HandledPrint(w, "ParseForm() error: "+err.Error())
		return
	}

	utils.HandledPrint(w, "Post Request Successful\n")
	printFormData(w, r)
}

// printFormData prints the form data to the response writer
func printFormData(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	address := r.FormValue("address")

	utils.HandledPrint(w, "Name: "+name+"\n")
	utils.HandledPrint(w, "Email: "+email+"\n")
	utils.HandledPrint(w, "Address: "+address+"\n")
}
