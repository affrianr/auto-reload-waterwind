package usecase

import (
	"github.com/affrianr/auto-reload-waterwind/rest-api/models"
)

type WaterWindUsecase interface {
	GetAll() ([]models.WaterWind, error)
	Create(waterWind *models.WaterWind) (*models.WaterWind, error)
	Update(id uint, waterWind *models.WaterWind) (*models.WaterWind, error)
	Delete(id uint) error
}
