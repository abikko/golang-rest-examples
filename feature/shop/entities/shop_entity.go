package entities

import (
	address "awesomeProject/feature/address/entities"
	position "awesomeProject/feature/position/entities"
	"time"
)

type (
	InsertShopDto struct {
		Id     uint32 `gorm:"primaryKey;autoIncrement" json:"id"`
		Amount uint32 `json:"amount"`
	}

	Shop struct {
		Id        uint32              `json:"id"`
		Amount    uint32              `json:"amount"`
		Name      string              `json:"name"`
		CreatedAt time.Time           `json:"createdAt"`
		Address   address.Address     `json:"address"`
		Positions []position.Position `json:"positions"`
	}
)
