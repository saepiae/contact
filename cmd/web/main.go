package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/saepiae/contact/pkg/models/postgres"

	"github.com/jackc/pgx"
	_ "github.com/jackc/pgx/v4"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	contacts *postgres.ContactModel
}

func main() {
	addr := flag.String("addr", ":8090", "Сетевой адрес HTTP")
	dbHost := flag.String("dbHost", "localhost:5432", "Хост подключения к БД")
	dbName := flag.String("dbName", "postgres", "Имя БД")
	dbUser := flag.String("dbUser", "user", "Имя пользователя при подключении к БД")
	dbPassword := flag.String("dbPasw", "password", "Пароль пользователя БД")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.LstdFlags)
	errorLog := log.New(os.Stderr, "ERROR\t", log.LstdFlags|log.Lshortfile)

	db, err := openDB(*dbHost, *dbName, *dbUser, *dbPassword)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	infoLog.Printf("Число соединений в пуле %d", db.Stat().MaxConnections)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		contacts: &postgres.ContactModel{ConnPool: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routers(),
	}

	infoLog.Printf("Запуск веб-сервера на %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dbHost string, dbName string, dbUser string, dbPasw string) (*pgx.ConnPool, error) {
	conf := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     dbHost,
			Database: dbName,
			User:     dbUser,
			Password: dbPasw,
		},
		MaxConnections: 5,
	}

	db, err := pgx.NewConnPool(conf)
	return db, err
}
