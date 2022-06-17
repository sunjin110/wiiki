#!/bin/sh

POSTGRES_MIGRATIONS_DIR=infra/postgres/migrations

echo "start migrate"
./goose -dir=$POSTGRES_MIGRATIONS_DIR postgres "host=${DB_HOST} port=${DB_PORT} user=${DB_USERNAME} dbname=${DB_NAME} sslmode=disable password=${DB_PASSWORD}" up
echo "finished migrate"
