createDb:
	docker exec -it postgres-container_name createDb --username=root --owner=root simple_bank

dropDb:
	docker exec -it postgres-container_name dropDb simple_bank

# postgres docker run -name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alphine
# mockgen -destination db/mock/store.go github.com/emmybritt/bank_app/db/sqlc Store
migrateup:
	migrate -path db/migration -database "postgresql://postgres:baped@localhost:5432/go_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:baped@localhost:5432/go_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test ./... -v -cover

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/emmybritt/bank_app/db/sqlc Store

.PHONY: postgres createDb dropDb migrateup migratedown sqlc test server mock


