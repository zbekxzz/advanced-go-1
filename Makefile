migrateup:
	migrate -path migrations -database "postgres://postgres:Zbekxzz3@localhost:5432/b.karakuzovDB?sslmode=disable" up
migratedown:
	migrate -path migrations -database "postgres://postgres:Zbekxzz3@localhost:5432/b.karakuzovDB?sslmode=disable" down
run:
	@go run ./cmd/api