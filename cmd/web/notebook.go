package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":8090", "Сетевой адрес HTTP")
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", home)
	mux.HandleFunc("/contact/all", allContacts)
	mux.HandleFunc("/contact/get", contact)
	mux.HandleFunc("/contact/create", createContact)
	mux.HandleFunc("/contact/edit", editContact)
	mux.HandleFunc("/contact/delete", deleteContact)
	mux.HandleFunc("/contact/dublicates", findDublicatedContacts)

	log.Printf("Запуск веб-сервера на %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
