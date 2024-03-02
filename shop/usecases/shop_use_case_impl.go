package usecases

import (
	"awesomeProject/shop/entities"
	"awesomeProject/shop/models"
	"awesomeProject/shop/repositories"
)

type shopUseCaseImpl struct {
	shopRepository repositories.ShopRepository
}

func NewShopUseCaseImpl(
	shopRepository repositories.ShopRepository,
) ShopUseCase {
	return &shopUseCaseImpl{
		shopRepository: shopRepository,
	}
}

func (u *shopUseCaseImpl) ShopDataProcessing(in *models.AddShopData) error {
	insertShopData := &entities.InsertShopDto{
		Amount: in.Amount,
	}

	if err := u.shopRepository.InsertShopData(insertShopData); err != nil {
		return err
	}

	return nil
}
