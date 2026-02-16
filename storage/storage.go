package storage

import (
	"3-struct/bin"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func Save(list bin.List) {
	jsonData, err := json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("data.json", jsonData, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

func Read() (bin.List, error) {
	var list bin.List
	fileDta, err := os.ReadFile("data.json")
	if err != nil {
		return nil, errors.New("файл для чтения не найден")
	}

	err = json.Unmarshal(fileDta, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
