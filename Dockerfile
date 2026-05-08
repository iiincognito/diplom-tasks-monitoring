# Базовый образ
FROM ubuntu:latest

# Установка необходимых пакетов
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Создание рабочей директории
WORKDIR /app

# Копирование исполняемого файла и web-файлов
COPY scheduler-server ./
COPY web ./web

# Порт веб-сервера
EXPOSE 7540

# Переменные окружения (можно переопределить при запуске)
ENV TODO_PORT=7540
ENV TODO_DBFILE=/data/scheduler.db
ENV TODO_PASSWORD=""

# Точка монтирования для базы данных
VOLUME ["/data"]

# Запуск сервера
CMD ["./scheduler-server"]
