migrateup:
	migrate -path migrations -database "postgres://postgres:Zbekxzz3@localhost:5432/b.karakuzovDB?sslmode=disable" -verbose up
migratedown:
	migrate -path migrations -database "postgres://postgres:Zbekxzz3@localhost:5432/b.karakuzovDB?sslmode=disable" -verbose down