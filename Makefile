setup:
	go install github.com/spf13/cobra-cli@latest
	go mod download
	go mod tidy