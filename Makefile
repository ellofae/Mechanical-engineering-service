.PHONY: postgres migrate

local_migrate_up:
	golang-migrate -path $(PWD)/platform/migrations -database "postgres://postgres:password@host.docker.internal/postgres?sslmode=disable" up

local_migrate_down:
	golang-migrate -path $(PWD)/platform/migrations -database "postgres://postgres:password@host.docker.internal/postgres?sslmode=disable" down

postgres:
	docker run --name postgres-db -e POSTGRES_PASSWORD=password -d --network host postgres

rebuild_compose:
	docker-compose up -d --no-deps --build

run:
	docker compose up