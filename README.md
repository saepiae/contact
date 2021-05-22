# contact
Простейший проект для изучения основ go.
Записаная книжка контактов
REST, postgres

***
Запуск из командной строки без флагов `go run ./cmd/web`

Справка по флагам `go run ./cmd/web -help`

Проверка суммы `go mod verify`

Загрузить все зависимости для проекта `go mod download`

Установка драйвера *github.com/go-sql-driver/mysql* последней версии `go get github.com/go-sql-driver/mysql`
Установка драйвера *github.com/go-sql-driver/mysql* с версией v1 `go get github.com/go-sql-driver/mysql@v1`

Обновить пакет *github.com/foo/bar* до последней доступной версии или исправления `go get -u github.com/foo/bar`
Обновить пакет *github.com/foo/bar* до версии *v2.0.0* `go get -u github.com/foo/bar@v2.0.0`

Удалить пакет *github.com/foo/bar* `go get github.com/foo/bar@none`
Удалить все неиспользуемые в проекте пакеты `go mod tidy` или `go mod tidy -v`