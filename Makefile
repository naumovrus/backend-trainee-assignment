build:
	docker-compose build backend-trainee-assignment

run:
	docker-compose up backend-trainee-assignment

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up

test:
	go test -v ./...

