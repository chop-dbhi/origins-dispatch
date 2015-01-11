all: test

clean:
	go clean ./...

doc:
	godoc -http=:6060

build: clean
	go build .

install: clean
	go install .

test: build
	go test -cover ./...

fmt:
	go vet ./...
	go fmt ./...

lint:
	golint ./...

.PHONY: test
