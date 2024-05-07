package entities

type UsersCredentials struct {
	Username string `json:"username" db:"username" form:"username" binding:"required" validate:"required,min=5,max=50"`
	Password string `json:"password" db:"password" form:"password" binding:"required" validate:"required,min=8"`
}
