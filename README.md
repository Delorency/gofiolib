## Как запустить
в следующем файле вписать необходимые параметры
```bash
    ./internal/configs/.env
```
### Через docker compose
Далее в корне проекта запустить
```bash
    docker compose --env-file=./internal/configs/.env up
```
сервис будет ожидать на localhost:${DOCKER_HOST_PORT}
### В терминале
Далее в корне проекта запустить
```bash
    go run ./cmd/gofiolib/main.go
```
сервис будет ожидать на localhost:${PORT}
