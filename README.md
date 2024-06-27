# Реализация итогового проекта по курсу "Go-разработчик с нуля" в Яндекс.Практикум

Проект представляет собой веб-сервер, напианный на Go и реализуещий функцтиональность планировщика задач (TODO-листа).
В проекте реализован функционал REST API, работа с базой данных и аутентификация по паролю.

## Выполненные задния со звездочкой

1. Определение порта сервера в env-файле
2. Определение пути к файлу базы данных в env-файле
3. Реализованы правила повторения задач:
   - повторение в указанные дни недели
   - повторение в указанные дни месяца
4. Реализована возможность поиска задач
5. Подключен механизм аутентификации
6. Подготовлен Dockerfile для создания образа

## Запуск проекта локально

1.  Для запуска проекта локально в директории выполнить комаду

```bash
   $ go run ./cmd/main.go
```

2.  Для компиляции проекта выполнить команду

```bash
    $ go build -o todo ./cmd/main.go
```

После компиляции запускать из директории проекта выполнив команду

```bash
   $ ./todo
```

3.  Для работы с TODO-листом в адресной строке браузера набрать адрес
    '''
    localhost:7540
    '''
4.  Для успешной работы программы в её директории должен находится env-файл со следующим содержимым:

```bash
WEB_DIR=./web
TODO_PORT=7540
TODO_DB_DIR=./db/
TODO_DBFILE=scheduler.db
TODO_DATE_LAYOUT=20060102
TODO_SEARCH_LAYOUT=02.01.2006
TODO_PASSWORD=12345
APP_LOG_DIR=./log/
APP_LOG_FILE=app.log
```

## Запуск тестов

В директории `tests` находятся тесты для проверки API, которое реализовано в веб-сервере.
Запуск тестов осуществляется командой

```bash
   $ go test ./tests
```

Настройки тестов находятся в файле `test/settings.go`.
Для успешного прохождения тестов он должен содержать следующие команды:

```go
var Port = 7540
var DBFile = "../db/scheduler.db"
var FullNextDate = true
var Search = true
var Token = `"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJoYXNoIjoiNTk5NDQ3MWFiYjAxMTEyYWZjYzE4MTU5ZjZjYzc0YjRmNTExYjk5ODA2ZGE1OWIzY2FmNWE5YzE3M2NhY2ZjNSJ9.EKlvvMlZ450BVUO__owlT3mkJ2NnhMIMr_OdXFzl95U"`
```

## Сборка и запуск проекта с помощью Docker

Для сборки докер-образа подготовлен Dockerfile. Файл должен содержать слудующие комнады:

```bash
    FROM golang:1.22.3

    EXPOSE 7540

    ENV WEB_DIR ./web/
    ENV TODO_PORT 7540
    ENV TODO_DB_DIR ./db/
    ENV TODO_DBFILE scheduler.db
    ENV TODO_DATE_LAYOUT 20060102
    ENV TODO_SEARCH_LAYOUT 02.01.2006
    ENV TODO_PASSWORD 12345
    ENV APP_LOG_DIR ./log/
    ENV APP_LOG_FILE app.log

    WORKDIR /usr/src/app

    COPY ./api ./api
    COPY ./cmd ./cmd
    COPY ./configs ./configs
    COPY ./model ./model
    COPY ./repository ./repository
    COPY ./service ./service
    COPY ./web ./web

    COPY go.mod go.sum ./

    RUN go mod download

    RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o todo ./cmd/main.go

    CMD ["./todo"]
```

Сборка образа выполняется командой

```bash
$ docker build -t plasmatrip/todo --tag todo:v1 .
```

Запуск готовго докер-образа выполняется командой

```bash
$ docker run -d -p 7540:7540 todo:v1
```

## Дополнительная информация

Директория `web` содержит файлы фронтенда для работы TODO-листа.
