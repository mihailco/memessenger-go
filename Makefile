build:
	docker-compose build  


up:
	docker-compose up

run:
	go run cmd/main.go

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:postgrespw@localhost:8000?sslmode=disable' up

swag:
	swag init -g cmd/main.go

