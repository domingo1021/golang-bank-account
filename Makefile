postgres:
	docker run --name bank_postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password1234 -d postgres:15-alpine
createdb:
	docker exec -it bank_postgres dropdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it bank_postgres dropdb --username=root --owner=root simple_bank
migrateup:
	migrate -path ./Docker/DB/migration -database "postgresql://root:password1234@localhost:5432/bank_account?sslmode=disable" -verbose up
migratedown:
	migrate -path ./Docker/DB/migration -database "postgresql://root:password1234@localhost:5432/bank_account?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown