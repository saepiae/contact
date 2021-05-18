package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Упс... что-то пошло не так", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Тут будет подробно описан контакт"))
}

func allContacts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Тут доступен список всех контактов"))
}

func createContact(w http.ResponseWriter, r *http.Request) {
	disabled, w := handlerAllowedMethod(w, r, http.MethodPost)
	if disabled {
		return
	}
	w.Write([]byte("Тут будет возможность добавить новый контакт"))
}

func editContact(w http.ResponseWriter, r *http.Request) {
	disabled, w := handlerAllowedMethod(w, r, http.MethodPut)
	if disabled {
		return
	}
	w.Write([]byte("Тут мы будем редактировать контакт"))
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	disabled, w := handlerAllowedMethod(w, r, http.MethodDelete)
	if disabled {
		return
	}
	w.Write([]byte("Удаляет существующий контакт"))
}

func findDublicatedContacts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Список дублирующихся контактов"))
}

func handlerAllowedMethod(w http.ResponseWriter, r *http.Request, method string) (bool, http.ResponseWriter) {
	forbidden := r.Method != method
	w.Header().Set("Content-Type", "application/json")
	if forbidden {
		w.Header().Add("Allow", method)
		w.WriteHeader(405)
		w.Write([]byte("Поддерживается только " + method + " метод"))
	}
	return forbidden, w
}
