package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":8090", "Сетевой адрес HTTP")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.LstdFlags)
	errorLog := log.New(os.Stderr, "ERROR\t", log.LstdFlags|log.Lshortfile)
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", home)
	mux.HandleFunc("/contact/all", allContacts)
	mux.HandleFunc("/contact/get", contact)
	mux.HandleFunc("/contact/create", createContact)
	mux.HandleFunc("/contact/edit", editContact)
	mux.HandleFunc("/contact/delete", deleteContact)
	mux.HandleFunc("/contact/dublicates", findDublicatedContacts)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Запуск веб-сервера на %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
