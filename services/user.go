package services

import (
	"GO-Boiler-Plate/models"
	"GO-Boiler-Plate/repositories"
)

// UserService Represented Constructor Function Of UserService
type UserService struct {
	repo repositories.UserRepository
}

// Get Represented of function to get the data from DB
func (service *UserService) Get() ([]models.User, error) {
	data, err := service.repo.Get()
	return data, err
}
