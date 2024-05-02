package dtos

type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	UpdateAt string `json:"update_at" db:"update_at"`
	CreateAt string `json:"create_at" db:"create_at"`
}
