## Как запустить
в следующих файлах вписать необходимые параметры
```bash
    ./internal/configs/.env
    ./internal/configs/.docker_env
```
### Через docker compose
Далее в корне проекта запустить
```bash
    docker compose --env-file=./internal/configs/.env --env-file=./internal/configs/.docker_env up -d
```
или
```bash
    make dockerup
```
сервис будет ожидать на localhost:${DOCKER_HOST_PORT}
### В терминале
Далее в корне проекта запустить
```bash
    go run ./cmd/gofiolib/main.go
```
или
```bash
    make terminalup
```
сервис будет ожидать на localhost:${PORT}
