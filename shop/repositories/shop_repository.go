package repositories

import "awesomeProject/shop/entities"

type ShopRepository interface {
	InsertShopData(in *entities.InsertShopDto) error
}
