package dto

import "ginEssential/model"

type UserDTO struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func ToUserDTO(user model.User) UserDTO {
	return UserDTO{
		Name:  user.Name,
		Phone: user.Phone,
	}
}
