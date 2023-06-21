sqlc: ## Execute SQL generation golang code
	docker run --rm -v $(pwd):/src -w /src kjconroy/sqlc generate

dev: ## Execute development
	go run cmd/main.go

up: ## Execute docker
	docker-compose up -d && docker exec -it app sh
