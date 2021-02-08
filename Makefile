
createdb:
		docker exec -it bookshelf-postgres createdb --username=docker --owner=docker bookshelf

dropdb:
		docker exec -it bookshelf-postgres dropdb bookshelf

migrateup:
		migrate -path app/database/migrations -database "postgresql://docker:docker@db:5432/bookshelf?sslmode=disable" -verbose up

migratedown:
		migrate -path app/database/migrations -database "postgresql://docker:docker@db:5432/bookshelf?sslmode=disable" -verbose down

sqlc:
		sqlc generate

.PHONY: createdb dropdb migrateup migratedown sqlc
