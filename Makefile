.DEFAULT_GOAl := run-cmd

.PHONY: fmt vet build run clean tidy

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build-cmd: vet
	go build -o gosc ./cmd/gosc

run-cmd: vet
	go run ./cmd/gosc/

clean:
	rm gosc && go clean ./... && go mod tidy
