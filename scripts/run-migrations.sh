!#bin/bash
migrate -path ./app/db/migration -database "postgresql://docker:docker@db:5432/bookshelf?sslmode=disable" -verbose up
