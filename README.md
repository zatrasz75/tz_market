# tz_market
Задание (Golang + PostgreSQL + Gin)

Написать сервис с тремя эндпоинтами:

1) принимает на вход Строение (Building) и записывает его в бд (поля Building: название, город, год сдачи, кол-во этажей)
2) возвращает список строений (Buildings), с возможностью фильтрации по городу, году и кол-ву этажей (параметры не обязательные)
3) OpenApi документация (генерировать из аннотаций например с помощью https://github.com/swaggo/swag)

Настройки соединения с Postgres читать из config файла:

host
port
user
password
db

## **Приступая к работе**

### **Предварительные требования**
- Перейти на версию Golang 1.22 или более позднюю
- PostgreSQL latest или более поздней версии

### **Установка**

**Клонировать репозиторий**\:

```
https://github.com/zatrasz75/tz_market.git
cd tz_song_libraries
```
- Необходимо отредактировать файл configs/configs.yml
```azure
server:
host: localhost
port: 8283
  cors-allowed-origins:
    - "https://api.example.com"
    - "https://example.com"
    - "http://localhost:8787"

database:
host: localhost
user: zatrasz
password: postgrespw
db: db-market
port: 49777
```
- Файл миграции находится в migrations
- Для создания новых файлов и запуска приложения используйте Makefile
```azure
run:
	go run cmd/main.go

up:
	sql-migrate new up

down:
	sql-migrate down

swag:
	swag init -d internal/handlers/ -g router.go --parseDependency --parseDepth 3
```
#### Запуск о умолчанию приложение:

```azure
[INFO] postgres.go:83 Применена 0 миграция!
[INFO] app.go:51 Запуск сервера на http://localhost:8283
[INFO] app.go:52 Документация Swagger API: http://localhost:8283/swagger/index.html
```

### Эндпоинты:
- /en/building [post] Создание новой записи строений
- /en/building [get] Список строений, с возможностью фильтрации по городу, году и кол-ву этажей(параметры не обязательные)
- Документация Swagger API: http://localhost:8283/swagger/index.html
