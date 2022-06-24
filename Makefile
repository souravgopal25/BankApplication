postgres:
	docker run --name postgresContainer -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=pass -d postgres:12-alpine

createdb:
	docker exec -it postgresContainer createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgresContainer drop simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:pass@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:pass@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb  -destination db/mock/store.go github.com/souravgopal25/BankApplication/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock