# Makefile для сборки Docker образа

BINARY_NAME=scheduler-server
DOCKER_IMAGE=scheduler:latest

.PHONY: all build docker clean

all: build docker

# Сборка бинарника для Linux (для Docker)
build:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) ./cmd/app/main.go

# Сборка бинарника для текущей ОС
build-local:
	go build -o $(BINARY_NAME) ./cmd/app/main.go

# Создание Docker образа
docker: build
	docker build -t $(DOCKER_IMAGE) .

# Запуск контейнера
docker-run:
	docker run -d \
		-p 7540:7540 \
		-e TODO_PORT=7540 \
		-e TODO_DBFILE=/data/scheduler.db \
		-e TODO_PASSWORD=$(TODO_PASSWORD) \
		-v $(PWD)/data:/data \
		--name scheduler \
		$(DOCKER_IMAGE)

# Запуск через docker-compose
up:
	mkdir -p data
	docker-compose up -d

# Остановка
stop:
	docker stop scheduler || true
	docker-compose down || true

# Очистка
clean:
	rm -f $(BINARY_NAME)
	docker rm -f scheduler || true
	docker rmi $(DOCKER_IMAGE) || true

# Просмотр логов
logs:
	docker logs -f scheduler
