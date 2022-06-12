MIGRATE_NAME=DEFAULT

gqlgen_generate:
	gqlgen generate

create_migrate:
	goose -path infra/postgres create $(MIGRATE_NAME)

goose_up:
	goose -path infra/postgres up

goose_status:
	goose -path infra/postgres status