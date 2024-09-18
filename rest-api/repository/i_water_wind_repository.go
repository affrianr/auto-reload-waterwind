package repository

import "github.com/affrianr/auto-reload-waterwind/rest-api/models"

type WaterWindRepository interface {
	GetAll() ([]models.WaterWind, error)
	Create(waterWind *models.WaterWind) error
	Update(id uint, waterWind *models.WaterWind) error
	Delete(id uint) error
}
