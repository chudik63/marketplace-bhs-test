up:
	docker-compose up -d

down:
	docker-compose down

dev-up:
	docker-compose up -d postgres

dev-down:
	docker-compose down postgres

run:
	go run cmd/main.go

migrate-up:
	migrate -source file://./migrations -database postgres://root:123@localhost:5432/marketplace?sslmode=disable up 2

migrate-down:
	migrate -source file://./migrations -database postgres://root:123@localhost:5432/marketplace?sslmode=disable down 2