build:
	docker-compose up --build

up:
	docker-compose up

down:
	docker-compose down

dev-up:
	docker-compose up postgres

dev-down:
	docker-compose down postgres