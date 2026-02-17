package main

import (
	"3-struct/bin"
	"3-struct/file"
	"3-struct/storage"
	"fmt"
)

func main() {
	bin1 := bin.NewBin("123", false, "example")
	bin2 := bin.NewBin("456", true, "test")

	binList := bin.NewBinList()
	binList = append(binList, bin1)
	binList = append(binList, bin2)

	files := file.NewOSFileReader()
	st := storage.NewJSONStorage(files, "data.json")

	if err := st.Save(binList); err != nil {
		fmt.Println("save error:", err)
		return
	}

	loaded, err := st.Read()
	if err != nil {
		fmt.Println("read error:", err)
		return
	}

	fmt.Println("loaded:", loaded)
}
