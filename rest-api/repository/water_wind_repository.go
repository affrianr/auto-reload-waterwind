package repository

import (
	"github.com/affrianr/auto-reload-waterwind/rest-api/models"
	"gorm.io/gorm"
)

type waterWindRepo struct {
	db *gorm.DB
}

func NewWaterWindRepository(db *gorm.DB) WaterWindRepository {
	return &waterWindRepo{db: db}
}

func (repo *waterWindRepo) GetAll() ([]models.WaterWind, error) {
	var data []models.WaterWind
	if err := repo.db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *waterWindRepo) Create(waterWind *models.WaterWind) error {
	return repo.db.Create(waterWind).Error
}

func (repo *waterWindRepo) Update(id uint, waterWind *models.WaterWind) error {
	return repo.db.Model(&models.WaterWind{}).Where("id = ?", id).Updates(waterWind).Error
}

func (repo *waterWindRepo) Delete(id uint) error {
	return repo.db.Delete(&models.WaterWind{}, id).Error
}
