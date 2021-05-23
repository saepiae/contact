package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/saepiae/contact/pkg/models"
)

type NewContact struct {
	FirstName  string         `json:"firstName"`
	LastName   string         `json:"lastName"`
	MiddleName sql.NullString `json:"middleName"`
	Phone      string         `json:"phone"`
	Email      sql.NullString `json:"email"`
	Address    sql.NullString `json:"address"`
}

type ShortContact struct {
	Id         int            `json:"id"`
	FirstName  string         `json:"firstName"`
	LastName   string         `json:"lastName"`
	MiddleName sql.NullString `json:"middleName"`
}

// Просто рандомная страница, которая возвращает некую html- страничку
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

// Возвращает контакт с указанным id
func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.notFound(w)
		return
	}
	contact, err := app.contacts.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	data, err := json.Marshal(contact)
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v\n", string(data))
}

// Возвращает все контакты
func (app *application) allContacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := app.contacts.FindAll()
	if err != nil {
		app.serverError(w, err)
		return
	}
	output := make(map[string]ShortContact)
	for index, value := range contacts {
		output[strconv.Itoa(index)] = ShortContact{value.ID, value.FirstName, value.LastName, value.MiddleName}
	}
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(output)
	if err != nil {
		app.serverError(w, err)
		return
	}
	fmt.Fprintf(w, "%v\n", string(data))
}

// Добавляет новый контакт
func (app *application) createContact(w http.ResponseWriter, r *http.Request) {
	disabled, w := handlerAllowedMethod(w, r, http.MethodPost, app)
	if disabled {
		return
	}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var contact NewContact
	err := decoder.Decode(&contact)
	if err != nil {
		app.clientError(w, 400)
		return
	}

	id, err := app.contacts.Insert(contact.FirstName, contact.LastName, contact.MiddleName, contact.Phone, contact.Email, contact.Address)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/contact/get?id=%d", id), http.StatusSeeOther)
}

// Редактирует существующий контакт
func (app *application) editContact(w http.ResponseWriter, r *http.Request) {
	disabled, w := handlerAllowedMethod(w, r, http.MethodPut, app)
	if disabled {
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var contact NewContact
	err = decoder.Decode(&contact)
	if err != nil {
		app.clientError(w, 400)
		return
	}

	id, err = app.contacts.Update(id, contact.FirstName, contact.LastName, contact.MiddleName, contact.Phone, contact.Email, contact.Address)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/contact/get?id=%d", id), http.StatusSeeOther)
}

// Удаляет контакт по указанному id
func (app *application) deleteContact(w http.ResponseWriter, r *http.Request) {
	disabled, w := handlerAllowedMethod(w, r, http.MethodDelete, app)
	if disabled {
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.notFound(w)
		return
	}
	_, err = app.contacts.Delete(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	http.Redirect(w, r, fmt.Sprintln("/contact/all"), http.StatusSeeOther)
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
