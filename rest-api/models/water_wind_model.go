package models

type WaterWind struct {
	ID    uint `gorm:"primary_key"`
	Water int  `json:"water"`
	Wind  int  `json:"wind"`
}
