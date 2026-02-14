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

func newBinList() BinList {
	return BinList{}
}

func newBin(id string, private bool, name string) Bin {
	return Bin{
		id:         id,
		private:    private,
		created_at: time.Now(),
		name:       name,
	}
}

func main() {
	bin1 := newBin("123", false, "example")
	bin2 := newBin("456", true, "test")

	binList := newBinList()
	binList = append(binList, bin1)
	binList = append(binList, bin2)

	fmt.Println(binList)
}
