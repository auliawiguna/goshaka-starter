build:
	go build -o server main.go

swagger:
	$(GOPATH)/bin/swag init --parseDependency --parseDepth 4 -g main.go --output docs/

critic:
	$(GOPATH)/bin/gocritic check -enableAll ./...

security:
	$(GOPATH)/bin/gosec ./...

lint:
	golangci-lint run ./...

run: build
	./server

watch:
	$(GOPATH)/bin/reflex -s -r '\.go$$' make run

unit_test:
	go test ./test/unit -v

api_test:
	go test ./test/api -v