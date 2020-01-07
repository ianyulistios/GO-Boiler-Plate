package tables

import "time"

//User is a model represented table schema
type User struct {
	ID        int       `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string    `gorm:"type:varchar(100)"`
	Email     string    `gorm:"type:varchar(100);unique_index"`
	Password  string    `gorm:"type:varchar(200)"`
	CreatedAt time.Time `gorm:"type:timestamp"`
}
