package usecase

import (
	"github.com/affrianr/auto-reload-waterwind/rest-api/models"
	"github.com/affrianr/auto-reload-waterwind/rest-api/repository"
)

type waterWindUsecase struct {
	waterWindRepo repository.WaterWindRepository
}

func NewWaterWindUsecase(repo repository.WaterWindRepository) WaterWindUsecase {
	return &waterWindUsecase{waterWindRepo: repo}
}

func (uc *waterWindUsecase) GetAll() ([]models.WaterWind, error) {
	data, err := uc.waterWindRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (uc *waterWindUsecase) Create(waterWind *models.WaterWind) (*models.WaterWind, error) {
	err := uc.waterWindRepo.Create(waterWind)
	return waterWind, err
}

func (uc *waterWindUsecase) Update(id uint, waterWind *models.WaterWind) (*models.WaterWind, error) {

	err := uc.waterWindRepo.Update(id, waterWind)
	return waterWind, err
}

func (uc *waterWindUsecase) Delete(id uint) error {
	return uc.waterWindRepo.Delete(id)
}
