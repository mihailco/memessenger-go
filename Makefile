build:
	docker-compose build  


up:
	docker-compose up


run:
	docker run memessenger -p 5432:5432

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:postgrespw@localhost:5436?sslmode=disable' up

swag:
	swag init -g cmd/main.go

