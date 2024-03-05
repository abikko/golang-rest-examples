package entities

type Position struct {
	Id     uint32 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name   string `json:"name"`
	Amount uint32 `json:"amount"`
	Price  uint32 `json:"price"`
}
