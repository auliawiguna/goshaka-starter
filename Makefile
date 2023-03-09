build:
	go build -o server -gcflags=all=-l -ldflags="-w -s" main.go

swagger:
	$(GOPATH)/bin/swag init --parseDependency --parseDepth 4 -g main.go --output docs/

critic:
	$(GOPATH)/bin/gocritic check -enableAll ./...

security:
	$(GOPATH)/bin/gosec ./...

lint:
	$(GOPATH)/bin/golangci-lint run ./...

run: build
	./server

watch:
	$(GOPATH)/bin/reflex -s -r '\.go$$' make run

unit_test:
	go test ./test/unit -v

api_test:
	go test ./test/api -v