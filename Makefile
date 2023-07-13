migrate_up:
	golang-migrate -path $(PWD)/platform/migrations -database "postgres://postgres:020603@localhost/postgres?sslmode=disable" up

migrate_down:
	golang-migrate -path $(PWD)/platform/migrations -database "postgres://postgres:020603@localhost/postgres?sslmode=disable" down