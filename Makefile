MIGRATE_BIN=$(HOME)/go/bin/sql-migrate

.PHONY: migrate-up migrate-down migrate-new migrate-status

migrate-up:
	@$(MIGRATE_BIN) up -config=dbconfig.yml -env=development

migrate-down:
	@$(MIGRATE_BIN) down -config=dbconfig.yml -env=development

migrate-new:
	@$(MIGRATE_BIN) new -config=dbconfig.yml -env=development $(name)

migrate-status:
	@$(MIGRATE_BIN) status -config=dbconfig.yml -env=development
