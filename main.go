package main

import (
	"fmt"
	"time"
)

type BinList []Bin
type Bin struct {
	id         string
	private    bool
	created_at time.Time
	name       string
}

func main() {
	bin1 := Bin{
		id:         "123",
		private:    false,
		created_at: time.Now(),
		name:       "example",
	}

	bin2 := Bin{
		id:         "456",
		private:    true,
		created_at: time.Now(),
		name:       "test",
	}

	binList := BinList{bin1}
	binList = append(binList, bin2)

	fmt.Println(binList)
}
