build:
	go build -o server main.go

swagger:
	swag init --parseDependency --parseDepth 4 -g main.go --output docs/

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./...

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run

unit_test:
	go test ./test/unit -v

api_test:
	go test ./test/api -v