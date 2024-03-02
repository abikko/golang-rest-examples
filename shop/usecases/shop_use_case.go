package usecases

import "awesomeProject/shop/models"

type ShopUseCase interface {
	ShopDataProcessing(in *models.AddShopData) error
}
