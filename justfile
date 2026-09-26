default: format test build

format:
  go fmt ./...

test:
  go test ./...

build:
  go build
