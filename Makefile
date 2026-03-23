include .env
export

MIGRATE_BIN=$(HOME)/go/bin/migrate
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

.PHONY: migrate-up migrate-down migrate-create migrate-force migrate-status

migrate-up:
	@$(MIGRATE_BIN) -path migrations -database "$(DB_URL)" up

migrate-down:
	@$(MIGRATE_BIN) -path migrations -database "$(DB_URL)" down 1

migrate-create:
	@$(MIGRATE_BIN) create -ext sql -dir migrations -seq $(name)

migrate-force:
	@$(MIGRATE_BIN) -path migrations -database "$(DB_URL)" force $(version)

migrate-status:
	@$(MIGRATE_BIN) -path migrations -database "$(DB_URL)" version
