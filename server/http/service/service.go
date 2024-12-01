package service

import (
	"server/models"
	"time"
)

func CreateNewService(s *models.Service) *models.Service {
	return &models.Service{
		ServiceId:   s.ServiceId,
		ServiceName: s.ServiceName,
		Price:       s.Price,
		CreatedAt:   time.Now(),
		ExpireAt:    time.Now().AddDate(0, 1, 0),
	}
}

func CalculateServiceTotal(services []models.Service) float32 {
	var total float32
	for _, service := range services {
		total += service.Price
	}
	return total
}
