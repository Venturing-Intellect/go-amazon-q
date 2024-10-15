run:
	go run main.go

mod-vendor:
	go mod vendor

linter:
	@golangci-lint run

gosec:
	@gosec -quiet ./...

validate: linter gosec

docker:
	docker-compose build
	docker-compose up

docker-run:
	docker run --env-file .env -p 8081:8081 -v $(pwd):/app travelis-backend:dev

migrate-create:
	@goose -dir=migrations create "$(name)" sql

migrate-up:
	@ export POSTGRES_HOST=localhost \
	&& export POSTGRES_USER=postgres \
	&& export POSTGRES_PASSWORD=mysecretpassword \
	&& export POSTGRES_DB=feedback_db \
	&& goose -dir=migrations postgres "host=${POSTGRES_HOST} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" up

migrate-down:
	@ export POSTGRES_HOST=localhost \
	&& export POSTGRES_USER=postgres \
	&& export POSTGRES_PASSWORD=mysecretpassword \
	&& export POSTGRES_DB=feedback_db \
	&& goose -dir=migrations postgres "host=${POSTGRES_HOST} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" down

swag:
	swag init -g main.go

install-swagger:
	go install github.com/swaggo/swag/cmd/swag@latest

get-swagdeps:
	go get github.com/swaggo/swag

swag-fmt:
	swag fmt