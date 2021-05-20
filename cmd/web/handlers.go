package main

import (
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Тут будет подробно описан контакт"))
}

func (app *application) allContacts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Тут доступен список всех контактов"))
}

func (app *application) createContact(w http.ResponseWriter, r *http.Request) {
	disabled, w := handlerAllowedMethod(w, r, http.MethodPost, app)
	if disabled {
		return
	}
	w.Write([]byte("Тут будет возможность добавить новый контакт"))
}

func (app *application) editContact(w http.ResponseWriter, r *http.Request) {
	disabled, w := handlerAllowedMethod(w, r, http.MethodPut, app)
	if disabled {
		return
	}
	w.Write([]byte("Тут мы будем редактировать контакт"))
}

func (app *application) deleteContact(w http.ResponseWriter, r *http.Request) {
	disabled, w := handlerAllowedMethod(w, r, http.MethodDelete, app)
	if disabled {
		return
	}
	w.Write([]byte("Удаляет существующий контакт"))
}

func findDublicatedContacts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Список дублирующихся контактов"))
}

func handlerAllowedMethod(w http.ResponseWriter, r *http.Request, method string, app *application) (bool, http.ResponseWriter) {
	forbidden := r.Method != method
	w.Header().Set("Content-Type", "application/json")
	if forbidden {
		w.Header().Add("Allow", method)
		app.clientError(w, http.StatusMethodNotAllowed)
	}
	return forbidden, w
}
