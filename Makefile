test:
	go test -race -v ./...

build:
	go build -race -i -o bin/storage