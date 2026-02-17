package file

import (
	"errors"
	"os"
	"path/filepath"
)

type FileReaderInterface interface {
	ReadFile(filename string) ([]byte, error)
	IsJsonFile(filename string) bool
}

func ReadFile(filename string) ([]byte, error) {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("файл для чтения не найден")
	}
	return fileData, nil
}

func IsJsonFile(filename string) bool {
	fileExt := filepath.Ext(filename)
	if fileExt == ".json" {
		return true
	}
	return false
}
