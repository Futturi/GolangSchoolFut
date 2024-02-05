sqlup:
	migrate -path internal/migrations -database "postgresql://root:12345@localhost:5432/root?sslmode=disable"  up
sqldown:
	migrate -path internal/migrations -database "postgresql://root:12345@localhost:5432/root?sslmode=disable"  down
run:
	go run cmd/main.go
.PHONY: sqlup sqldown run