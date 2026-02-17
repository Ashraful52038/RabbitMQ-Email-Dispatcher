.PHONY: up down publisher consumer clean

up:
	docker compose up -d
	@echo "================================="
	@echo "RabbitMQ UI: http://localhost:15673 (guest/guest)"
	@echo "Mailpit UI: http://localhost:8025"
	@echo "================================="

down:
	docker compose down

publisher:
	go run publisher/main.go

consumer:
	go run consumer/main.go

clean:
	docker compose down -v