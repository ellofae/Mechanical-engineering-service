.PHONY: postgres mongo migrate

local_migrate_up:
	golang-migrate -path $(PWD)/platform/migrations -database "postgres://postgres:password@host.docker.internal/postgres?sslmode=disable" up

local_migrate_down:
	golang-migrate -path $(PWD)/platform/migrations -database "postgres://postgres:password@host.docker.internal/postgres?sslmode=disable" down

mongo:
	docker run --name auth-mongo -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=example mongo

postgres:
	docker run --name postgres-db -e POSTGRES_PASSWORD=password -d --network host postgres

rebuild_compose:
	docker-compose up -d --no-deps --build

run:
	docker compose up