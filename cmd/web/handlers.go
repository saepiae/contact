package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/saepiae/contact/pkg/models"
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
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.notFound(w)
		return
	}
	c, err := app.contacts.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	w.Write([]byte("Тут будет подробно описан контакт c id = " + strconv.Itoa(id)))
	fmt.Fprintf(w, "%v", c)
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
	firstName := "firstName"
	lastName := "lastName"
	middleName := "middleName"
	phone := "phone"
	email := "email"
	address := "address"
	id, err := app.contacts.Insert(firstName, lastName, middleName, phone, email, address)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/contact/get?id=%d", id), http.StatusSeeOther)
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
