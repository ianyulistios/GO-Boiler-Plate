package repositories

import (
	"GO-Boiler-Plate/models"

	"github.com/jinzhu/gorm"
)

//CategoryRepository Represented of Sturct to call the gorm and connect to all func
type CategoryRepository struct {
	db *gorm.DB
}

//Get Represented of function to get the data of user
func (repo *CategoryRepository) Get() []models.User {
	userData := make([]models.User, 0)

	repo.db.Find(&userData)

	return userData
}
