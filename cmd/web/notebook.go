package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":8090", "Сетевой адрес HTTP")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.LstdFlags)
	errorLog := log.New(os.Stderr, "ERROR\t", log.LstdFlags|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", app.home)
	mux.HandleFunc("/contact/all", app.allContacts)
	mux.HandleFunc("/contact/get", app.contact)
	mux.HandleFunc("/contact/create", app.createContact)
	mux.HandleFunc("/contact/edit", app.editContact)
	mux.HandleFunc("/contact/delete", app.deleteContact)
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
