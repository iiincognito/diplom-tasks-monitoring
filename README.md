# Go Планировщик Задач (Task Scheduler)

Веб-приложение для управления задачами с поддержкой повторяющихся событий. Позволяет создавать, редактировать, удалять задачи и отмечать их выполнение.

## Функциональность

- ✅ Создание задач с указанием даты, заголовка, комментария и правила повторения
- ✅ Просмотр списка задач с сортировкой по дате
- ✅ Поиск задач по тексту или конкретной дате
- ✅ Редактирование задач
- ✅ Удаление задач
- ✅ Отметка задач выполненными (с автоматическим переносом повторяющихся задач на следующую дату)
- ✅ Поддержка различных правил повторения: ежедневно (d N), ежегодно (y), по дням недели (w 1,2,3), по месяцам (m 1,2 3,4)
- ✅ Аутентификация по паролю с JWT-токенами
- ✅ SQLite база данных

## Выполненные задания со звёздочкой

- ✅ **Аутентификация** — реализована через JWT-токены, пароль задаётся в `TODO_PASSWORD`
- ✅ **Docker** — созданы Dockerfile и docker-compose.yml для контейнеризации

## Запуск локально

### Требования

- Go 1.25+
- SQLite3

### Установка зависимостей

```bash
go mod tidy
```

### Запуск сервера

```bash
# Без пароля (открытый доступ)
go run ./cmd/app/main.go

# С паролем (требуется аутентификация)
TODO_PASSWORD=12345 go run ./cmd/app/main.go

# С кастомным портом и базой данных
TODO_PORT=8080 TODO_DBFILE=./mydb.db go run ./cmd/app/main.go
```

### Переменные окружения

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `TODO_PORT` | Порт веб-сервера | `7540` |
| `TODO_DBFILE` | Путь к файлу SQLite | `./scheduler.db` |
| `TODO_PASSWORD` | Пароль для аутентификации (пустой = без пароля) | `""` |
| `TODO_SECRET` | Секретный ключ для JWT | авто-генерация |

### Доступ в браузере

После запуска откройте: http://localhost:7540

Если установлен пароль, сначала нужно войти на странице `/login.html`.

## Запуск тестов

### Настройка tests/settings.go

```go
var Port = 7540              // Порт сервера
var DBFile = "../scheduler.db"  // Путь к базе данных
var FullNextDate = true      // Включить полные тесты правил повторения
var Search = true            // Включить тесты поиска
var Token = ""               // JWT токен (заполнить если есть TODO_PASSWORD)
```

### Получение токена (если установлен пароль)

```bash
# Запустите сервер с паролем
TODO_PASSWORD=12345 go run ./cmd/app/main.go

# Получите токен
curl -X POST http://localhost:7540/api/signin \
  -H "Content-Type: application/json" \
  -d '{"password":"12345"}'
```

Скопируйте токен из ответа в `var Token` в `tests/settings.go`.

### Запуск тестов

```bash
# Все тесты
go test ./tests

# Конкретный тест
go test -run ^TestApp$ ./tests

# С подробным выводом
go test -v ./tests
```

**Важно:** Сервер должен быть запущен перед запуском тестов!

## Docker

### Быстрый старт

```bash
# Многоэтапная сборка (рекомендуется)
docker build -f Dockerfile.multistage -t scheduler:latest .

# Запуск без пароля
docker run -d -p 7540:7540 -v $(pwd)/data:/data --name scheduler scheduler:latest

# Запуск с паролем
docker run -d -p 7540:7540 -e TODO_PASSWORD=12345 -v $(pwd)/data:/data --name scheduler scheduler:latest
```

### Docker Compose

```bash
# Сборка и запуск
docker-compose up -d

# Остановка
docker-compose down
```

### Makefile

```bash
# Сборка и запуск
make up

# Остановка
make stop

# Просмотр логов
make logs
```

Подробнее в [DOCKER.md](./DOCKER.md)

## API Endpoints

### Публичные (без аутентификации)
- `POST /api/signin` — вход, получение JWT токена
- `GET /api/nextdate` — расчёт следующей даты

### Защищённые (требуют аутентификации если `TODO_PASSWORD` установлен)
- `GET /api/tasks?search=` — список задач с поиском
- `GET /api/task?id=` — получить задачу по ID
- `POST /api/task` — создать задачу
- `PUT /api/task` — обновить задачу
- `DELETE /api/task?id=` — удалить задачу
- `POST /api/task/done?id=` — отметить выполненной

## Структура проекта

```
.
├── cmd/app/main.go              # Точка входа
├── internal/
│   ├── core/
│   │   ├── db/                  # Подключение к SQLite
│   │   ├── domain/              # Модели данных
│   │   └── transport/http/server/  # HTTP сервер
│   └── features/
│       ├── auth/                # Аутентификация
│       │   ├── service/
│       │   ├── middleware/
│       │   └── transport/http/
│       └── task/                # Задачи
│           ├── repository/      # Работа с БД
│           ├── service/         # Бизнес-логика
│           └── transport/http/  # HTTP обработчики
├── tests/                       # Интеграционные тесты
├── web/                         # Фронтенд
├── Dockerfile                   # Docker образ
├── Dockerfile.multistage        # Многоэтапная сборка
├── docker-compose.yml           # Docker Compose
└── README.md                    # Этот файл
```

## Лицензия

Учебный проект для курса Go-разработки.