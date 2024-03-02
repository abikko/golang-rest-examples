package entities

import "time"

type (
	InsertShopDto struct {
		Id        uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
		Amount    uint32    `json:"amount"`
		CreatedAt time.Time `json:"createdAt"`
	}

	Shop struct {
		Id        uint32    `json:"id"`
		Amount    uint32    `json:"amount"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"createdAt"`
	}
)
