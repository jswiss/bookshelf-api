
postgres:
	docker run --name db -p 5432:5432 -e POSTGRES_USER=docker -e POSTGRES_PASSWORD=docker -d postgres:13-alpine

test:
	go test -v -cover ./...

server:
	go run main.go

createdb:
		docker exec -it db createdb --username=docker --owner=docker bookshelf

dropdb:
		docker exec -it db dropdb bookshelf

migrateup:
		migrate -path app/database/migrations -database "postgresql://docker:docker@db:5432/bookshelf?sslmode=disable" -verbose up

migratedown:
		migrate -path app/database/migrations -database "postgresql://docker:docker@db:5432/bookshelf?sslmode=disable" -verbose down

sqlc:
		sqlc generate

server:
		go run main.go

up:
		@echo "=============starting api locally============="
		docker-compose up -d

logs:
		docker-compose logs -f

down:
		docker-compose down

test:
		go test -v -cover ./...

clean: down
		@echo "=============cleaning up============="
		rm -f api
		docker system prune -f
		docker volume prune -f

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server
