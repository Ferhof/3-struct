package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	fileData, err := os.ReadFile("data.json")
	if err != nil {
		return nil, errors.New("файл для чтения не найден")
	}

	return fileData, nil
}

func IsJsonFile(filename string) bool {
	fileData, err := os.Open(filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return false
	}
	var data interface{}
	decoder := json.NewDecoder(fileData)
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("JSON некорректен:", err)
		return false
	}
	return true
}
