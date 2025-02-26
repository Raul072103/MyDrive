include .env
MIGRATIONS_PATH = ./cmd/migrate/migrations

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir ${MIGRATIONS_PATH} $(word 2, $(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path $(MIGRATIONS_PATH) -database $(DB_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path $(MIGRATIONS_PATH) -database $(DB_ADDR) down $(word 2, $(MAKECMDGOALS))

.PHONY: seed
seed:
	@go run cmd/migrate/seed/main.go

.PHONY: gen-docs
gen-docs:
	@swag init --generalInfo api/main.go --dir cmd,internal/repo,internal/db,internal/env && swag fmt

.PHONY: test
test:
	@go test -v ./...

.PHONY: run-vite
run-vite:
	@npm --prefix ./web run build && npm --prefix ./web run dev