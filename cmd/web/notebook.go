package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", home)
	mux.HandleFunc("/contact/all", allContacts)
	mux.HandleFunc("/contact/get", contact)
	mux.HandleFunc("/contact/create", createContact)
	mux.HandleFunc("/contact/edit", editContact)
	mux.HandleFunc("/contact/delete", deleteContact)
	mux.HandleFunc("/contact/dublicates", findDublicatedContacts)

	port := 8090
	log.Println("Запуск веб-сервера на http://127.0.0.1:", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), mux)
	log.Fatal(err)
}
