package main

import (
	"log"
	"net/http"
	"strconv"
)

// Создается функция-обработчик "home", которая записывает байтовый слайс, содержащий
// текст "Привет из Snippetbox" как тело ответа.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет из Notebook"))
}

// Отображает один контакт
func contact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Тут будет подробно описан контакт"))
}

// Список всех контактов
func allContacts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Тут доступен список всех контактов"))
}

// Создать новый контакт
func createContact(w http.ResponseWriter, r *http.Request) {
	disabled, w := handlerAllowedMethod(w, r, http.MethodPost)
	if disabled {
		return
	}
	w.Write([]byte("Тут будет возможность добавить новый контакт"))
}

// Редактирование существующего контакта
func editContact(w http.ResponseWriter, r *http.Request) {
	disabled, w := handlerAllowedMethod(w, r, http.MethodPost)
	if disabled {
		return
	}
	w.Write([]byte("Тут мы будем редактировать контакт"))
}

// Удаление существующего контакта
func deleteContact(w http.ResponseWriter, r *http.Request) {
	disabled, w := handlerAllowedMethod(w, r, http.MethodPost)
	if disabled {
		return
	}
	w.Write([]byte("Удаляет существующий контакт"))
}

// Ищет дублирующие контакты
func findDublicatedContacts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Список дублирующихся контактов"))
}

// Проверяет поддержку http-метода
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
