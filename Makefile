.EXPORT_ALL_VARIABLES:
.ONESHELL:

setup:
	go install github.com/spf13/cobra-cli@latest
	go mod download
	go mod tidy

test:
	go test -v ./...

# TODO: Inject commit/tag info
build:
	mkdir -p ./bin
	rm -f ./bin/*
	GOOS=darwin; GOARCH=amd64;  go build -o bin/templar-$$GOOS-$$GOARCH
	GOOS=darwin; GOARCH=arm64;  go build -o bin/templar-$$GOOS-$$GOARCH
	GOOS=linux; GOARCH=amd64;   go build -o bin/templar-$$GOOS-$$GOARCH
	GOOS=windows; GOARCH=amd64; go build -o bin/templar-$$GOOS-$$GOARCH.exe