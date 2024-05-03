package entities

import "github.com/textures1245/go-template/internal/user/dtos"

func NewUserDetail(dat *User) *dtos.UserDetailRespond {
	return &dtos.UserDetailRespond{
		ID:          dat.ID,
		Name:        dat.Name,
		Email:       dat.Email,
		PhoneNumber: dat.PhoneNumber,
		IdCard:      dat.IdCard,
		CreatedAt:   dat.CreatedAt,
		UpdatedAt:   dat.UpdatedAt,
	}
}
