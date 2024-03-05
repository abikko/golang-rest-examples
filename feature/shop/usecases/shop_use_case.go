package usecases

import (
	"awesomeProject/feature/shop/models"
)

type ShopUseCase interface {
	ShopDataProcessing(in *models.AddShopData) error
}
