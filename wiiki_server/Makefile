MIGRATE_NAME=DEFAULT
POSTGRES_DEVELOP="host=127.0.0.1 port=5432 user=wiiki_user dbname=wiiki sslmode=disable password=ZnfZxXY3"
POSTGRES_MIGRATIONS_DIR=infra/postgres/migrations

gqlgen_generate:
	gqlgen generate

create_migrate:
	# TODO name
	./goose -dir=$(POSTGRES_MIGRATIONS_DIR) postgres $(POSTGRES_DEVELOP) create migrate go

goose_up:
	./goose -dir=$(POSTGRES_MIGRATIONS_DIR) postgres $(POSTGRES_DEVELOP) up

goose_down:
	./goose -dir=$(POSTGRES_MIGRATIONS_DIR) postgres $(POSTGRES_DEVELOP) down

goose_status:
	./goose postgres $(POSTGRES_DEVELOP) status

goose_build:
	go build -o goose cmd/migration_postgres/main.go