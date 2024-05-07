package entities

type UserUpdateReq struct {
	Name        string `json:"name" form:"name" binding:"required"  validate:"required,min=5,max=50"`
	Email       string `json:"email" form:"email" binding:"required" validate:"required,email"`
	PhoneNumber string `json:"phone_number" form:"phone_number" binding:"required" validate:"required,min=10,max=10"`
	IdCard      string `json:"id_card" form:"id_card" binding:"required" validate:"required,min=14,max=14"`
}
