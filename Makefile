postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root my_portfolio

dropdb:
	docker exec -it postgres12 dropdb my_portfolio

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/my_portfolio?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/my_portfolio?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go
	
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/gaku101/my-portfolio/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown mock