package objects

//UserRequest Represented of User Request
type UserRequest struct {
	Email    string `json:"email" form:"email" query:"email" validate:"required"`
	Password string `json:"password" form:"password" query:"password" validate:"required"`
}
