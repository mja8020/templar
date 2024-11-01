.EXPORT_ALL_VARIABLES:
.ONESHELL:

setup:
	go install gotest.tools/gotestsum@latest
	go install github.com/spf13/cobra-cli@latest
	go mod download
	go mod tidy

test:
	gotestsum --format testname --junitfile results.xml ./... -- -v

# TODO: Inject commit/tag info
build:
	mkdir -p ./bin
	rm -f ./bin/*
	GOOS=darwin; GOARCH=amd64;  go build -o bin/templar-$$GOOS-$$GOARCH
	GOOS=darwin; GOARCH=arm64;  go build -o bin/templar-$$GOOS-$$GOARCH
	GOOS=linux; GOARCH=amd64;   go build -o bin/templar-$$GOOS-$$GOARCH
	GOOS=windows; GOARCH=amd64; go build -o bin/templar-$$GOOS-$$GOARCH.exe
	GOOS=windows; GOARCH=arm64; go build -o bin/templar-$$GOOS-$$GOARCH.exe