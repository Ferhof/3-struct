package main

import (
	"3-struct/bin"
	"fmt"
)

type Vault interface {
	Read() string
}

func main() {
	bin1 := bin.NewBin("123", false, "example")
	bin2 := bin.NewBin("456", true, "test")

	binList := bin.NewBinList()
	binList = append(binList, bin1)
	binList = append(binList, bin2)

	fmt.Println(binList)
}
