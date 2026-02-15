package bin

import "time"

type List []Bin
type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func NewBinList() List {
	return List{}
}

func NewBin(id string, private bool, name string) Bin {
	return Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}
