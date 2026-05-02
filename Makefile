.PHONY: db-up db-down run-local docker-up docker-down

db-up:
	docker compose up -d db

db-down:
	docker compose stop db

run-local:
	go run back/cmd/api/main.go

docker-up:
	docker compose up --build -d

docker-down:
	docker compose down
