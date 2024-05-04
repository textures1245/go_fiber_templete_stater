package entities

type UserUpdateReq struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" binding:"required"`
	IdCard      string `json:"id_card" form:"id_card" binding:"required"`
}
