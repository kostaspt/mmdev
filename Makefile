CMD_DIRS=go list -f '{{.Dir}}' ./cmd/... | grep -v /vendor/
DIRS=go list -f '{{.Dir}}' ./... | grep -v /vendor/

.PHONY: build
build: clean
	mkdir -p bin/ && go build -ldflags "-X main.Version=${VERSION}" -o ./bin/ `$(call CMD_DIRS)`

.PHONY: clean
clean:
	go clean
	rm -rf ./bin

.PHONY: deps
deps:
	go mod tidy -v
	go mod verify

.PHONY: format
format:
	gofmt -w `$(call DIRS)`
