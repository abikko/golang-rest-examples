package repositories

import (
	"awesomeProject/shop/entities"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type shopPostgresRepository struct {
	db *gorm.DB
}

func NewShopPostgresRepository(db *gorm.DB) ShopRepository {
	return &shopPostgresRepository{db: db}
}

func (r *shopPostgresRepository) InsertShopData(in *entities.InsertShopDto) error {
	data := &entities.Shop{
		Amount: in.Amount,
	}

	result := r.db.Create(data)

	if result.Error != nil {
		log.Errorf("InsertShopData: %v", result.Error)
		return result.Error
	}

	log.Debugf("InsertShopData: %v", result.RowsAffected)
	return nil
}
