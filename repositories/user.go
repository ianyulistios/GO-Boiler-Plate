package repositories

import (
	"GO-Boiler-Plate/models"

	"github.com/jinzhu/gorm"
)

//UserRepository Represented of Sturct to call the gorm and connect to all func
type UserRepository struct {
	db *gorm.DB
}

//Get Represented of function to get the data of user
func (repo *UserRepository) Get() ([]models.User, error) {
	userData := make([]models.User, 0)

	err := repo.db.Find(&userData)

	return userData, err.Error
}
