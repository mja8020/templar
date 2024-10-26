package utils

import (
	"os"
)

func FileRead(path string) (content string, err error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return
	}

	content = string(bytes)

	return
}

func FileWrite(path string, content string) (err error) {
	exists, err := FileExists(path)
	if err != nil {
		return
	}

	var file *os.File

	if exists {
		file, err = os.Open(path)
	} else {
		file, err = os.Create(path)
	}

	if err != nil {
		return
	}

	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return
	}

	return
}

func FileExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
