package utils

import "os"

func ReadFile(path string) (content string, err error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return
	}

	content = string(bytes)

	return
}
