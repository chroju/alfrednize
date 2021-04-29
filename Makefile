BINARY_NAME=alfrednize

.PHONY: install test lint crossbuild clean

install:
	go install

lint:
	go mod tidy
	gofmt -s -l .
	golint ./...
	go vet ./...

test: lint
	go test -v ./...

crossbuild: test
	gox -os="linux darwin windows" -arch="386 amd64" -output "bin/remo_{{.OS}}_{{.Arch}}/{{.Dir}}"

mod:
	go mod download

build:
	go build -o bin/alfrednize

clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -f bin/

test-coverage: mod
	go test -race -covermode atomic -coverprofile=covprofile ./...
