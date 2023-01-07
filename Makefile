build:
	go build -o server main.go

swagger:
	swag init --parseDependency --parseDepth 4 -g main.go --output docs/

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run