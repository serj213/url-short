.PHONY: dc, run, test, coverage

dc:
	docker-compose up --remove-orphans --build

run:
	configPath=config/local.yaml go run cmd/main.go

test:
	go test ./... -coverprofile=coverage.out

coverage:
	go tool cover -func=coverage.out