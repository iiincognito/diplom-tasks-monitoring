# Docker

## Сборка образа

### Вариант 1: Многоэтапная сборка (рекомендуется)

Самый простой способ — использовать `Dockerfile.multistage`, который сам собирает бинарник:

```bash
# Сборка (не требует предварительной сборки бинарника)
docker build -f Dockerfile.multistage -t scheduler:latest .

# Запуск
docker run -d -p 7540:7540 -v $(pwd)/data:/data --name scheduler scheduler:latest
```

### Вариант 2: Через Makefile

```bash
# Сборка бинарника и Docker образа
make build docker

# Запуск контейнера
make docker-run

# Остановка
make stop
```

### Вариант 3: Через Docker Compose

```bash
# Сборка и запуск
make up

# Или вручную:
docker-compose up -d
```

### Вариант 4: Ручная сборка (требует предварительно собранный бинарник)

```bash
# 1. Собрать бинарник для Linux
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o scheduler-server ./cmd/app/main.go

# 2. Создать Docker образ
docker build -t scheduler:latest .

# 3. Запустить контейнер
docker run -d \
  -p 7540:7540 \
  -e TODO_PORT=7540 \
  -e TODO_DBFILE=/data/scheduler.db \
  -e TODO_PASSWORD=12345 \
  -v $(pwd)/data:/data \
  --name scheduler \
  scheduler:latest
```

## Переменные окружения

- `TODO_PORT` — порт веб-сервера (по умолчанию 7540)
- `TODO_DBFILE` — путь к базе данных SQLite (внутри контейнера `/data/scheduler.db`)
- `TODO_PASSWORD` — пароль для аутентификации (пустой = без пароля)
- `TODO_SECRET` — секретный ключ для JWT (опционально)

## Параметры запуска

При запуске контейнера важно:

1. **Порт** — пробросить порт 7540 (или указанный в `TODO_PORT`)
2. **Volume** — смонтировать директорию для базы данных, чтобы данные сохранялись между перезапусками контейнера

### Примеры

**Без пароля:**
```bash
docker run -d -p 7540:7540 -v $(pwd)/data:/data scheduler:latest
```

**С паролем:**
```bash
docker run -d \
  -p 7540:7540 \
  -e TODO_PASSWORD=mysecretpassword \
  -v $(pwd)/data:/data \
  scheduler:latest
```

**С кастомным портом:**
```bash
docker run -d \
  -p 8080:8080 \
  -e TODO_PORT=8080 \
  -v $(pwd)/data:/data \
  scheduler:latest
```

## Доступ к приложению

После запуска открой в браузере:
- http://localhost:7540 (или твой порт)

Если пароль установлен, сначала нужно войти на странице `/login.html`.
