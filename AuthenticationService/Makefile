DB_MIGRATION_FOLDER = database/migrations
run:
	go run .

migration-create:
	goose -dir ${DB_MIGRATION_FOLDER} create ${name} sql

migration-up:
	goose -dir ${DB_MIGRATION_FOLDER} mysql "${url}" up

build:
	go build -o app
