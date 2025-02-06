.PHONY: dc, run

dc:
	docker-compose up --remove-orphans --build

run:
	configPath=config/local.yaml go run cmd/main.go