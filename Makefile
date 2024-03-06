# docker-compose 사용 안할 경우에 사용.
postgres:
	docker run --name postgres14 -p 54321:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root voyager_db

dropdb:
	docker exec -it postgres14 dropdb voyager_db

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:54321/voyager_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:54321/voyager_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go clean -testcache; go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination ./db/mock/store.go github.com/yangoneseok/voyager/db/sqlc Store

# newmigrate:
# 	migrate create -ext sql -dir db/migration -seq <name>

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock
