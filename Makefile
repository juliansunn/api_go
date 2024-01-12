include .env
export

postgres:
	docker run --name postgres12 \
	-p $(DB_PORT):5432 \
	-e POSTGRES_USER=$(DB_USER) \
	-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
	-e POSTGRES_DB=$(DB_NAME) \
	-d postgres:12-alpine

migrateup:
	migrate -path db/migration -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)" -verbose up

migratecreate:
	migrate create -ext sql -dir db/migration -seq init_schema

setup:
	make postgres &&  make migratecreate

sqlc:
	sqlc generate

test:
	go test ./...

run:
	go run main.go

build:
	go build

.PHONY: sqlc test run build postgres createdb migrateup migreatecreate