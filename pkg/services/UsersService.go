package services

import "go.mod/pkg/entity"

type UserService interface {
	GetUsers() (string, int, string)
}

// constractor
func NewUsers(nama string, umur int, alamat string) UserService {
	var user = entity.Users{
		Name:    nama,
		Age:     umur,
		Address: alamat,
	}

	return user
}
