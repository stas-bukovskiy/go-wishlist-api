go-local-run:
	go run cmd/main.go

local-run:
	docker-compose -f docker-compose.local.yml up -d

docker-run:
	docker-compose up -d