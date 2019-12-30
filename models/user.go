package models

//User Represented of user model based on database
type User struct {
	Name     string `json:"name" form:"name" query:"name" validate:"required"`
	Email    string `json:"email" form:"email" query:"email" validate:"required"`
	Password string `json:"password" form:"password" query:"password" validate:"required"`
}
