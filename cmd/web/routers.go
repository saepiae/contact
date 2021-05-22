package main

import "net/http"

func (app *application) routers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", app.home)
	mux.HandleFunc("/contact/all", app.allContacts)
	mux.HandleFunc("/contact/get", app.contact)
	mux.HandleFunc("/contact/create", app.createContact)
	mux.HandleFunc("/contact/edit", app.editContact)
	mux.HandleFunc("/contact/delete", app.deleteContact)
	mux.HandleFunc("/contact/dublicates", findDublicatedContacts)
	return mux
}
