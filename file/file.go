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

type OSFileReader struct{}

func NewOSFileReader() *OSFileReader {
	return &OSFileReader{}
}

func (r *OSFileReader) ReadFile(filename string) ([]byte, error) {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("файл для чтения не найден")
	}
	return fileData, nil
}

func (r *OSFileReader) IsJsonFile(filename string) bool {
	return filepath.Ext(filename) == ".json"
}
