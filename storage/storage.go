package storage

import (
	"3-struct/bin"
	"3-struct/file"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Storage interface {
	Read() (bin.List, error)
	Save(list bin.List) error
}

type JSONStorage struct {
	files file.FileReaderInterface
	path  string
}

func NewJSONStorage(files file.FileReaderInterface, path string) *JSONStorage {
	return &JSONStorage{
		files: files,
		path:  path,
	}
}

func (s *JSONStorage) Save(list bin.List) error {
	if !s.files.IsJsonFile(s.path) {
		return errors.New("storage поддерживает только .json файл")
	}

	jsonData, err := json.Marshal(list)
	if err != nil {
		return err
	}

	if err := os.WriteFile(s.path, jsonData, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func (s *JSONStorage) Read() (bin.List, error) {
	if !s.files.IsJsonFile(s.path) {
		return nil, errors.New("storage поддерживает только .json файл")
	}

	var list bin.List

	fileData, err := s.files.ReadFile(s.path)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(fileData, &list); err != nil {
		return nil, err
	}

	return list, nil
}

// (необязательно) если где-то в проекте уже вызывали storage.Save(...) и storage.Read(),
// можно временно оставить совместимость через обёртки:
func Save(list bin.List) {
	fmt.Println("WARNING: используйте JSONStorage.Save(list) вместо storage.Save(list)")
	_ = os.WriteFile("data.json", mustMarshal(list), os.ModePerm)
}

func Read() (bin.List, error) {
	fmt.Println("WARNING: используйте JSONStorage.Read() вместо storage.Read()")
	var list bin.List
	b, err := os.ReadFile("data.json")
	if err != nil {
		return nil, errors.New("файл для чтения не найден")
	}
	if err := json.Unmarshal(b, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func mustMarshal(list bin.List) []byte {
	b, _ := json.Marshal(list)
	return b
}
