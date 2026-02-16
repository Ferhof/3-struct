package bin

import "time"

type List []Bin
type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

func NewBinList() List {
	return List{}
}

func NewBin(id string, private bool, name string) Bin {
	return Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}
