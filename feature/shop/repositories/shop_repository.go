package repositories

import (
	"awesomeProject/feature/shop/entities"
)

type ShopRepository interface {
	InsertShopData(in *entities.InsertShopDto) error
}
