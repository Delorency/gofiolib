dockerup:
	docker compose --env-file=./internal/configs/.env --env-file=./internal/configs/.docker_env up -d

terminalup:
	go run ./cmd/gofiolib/main.go
