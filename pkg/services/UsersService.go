package services

import "go.mod/pkg/entity"

type UserService interface {
	GetUsers() (string, int, string)
}

// constarctor
func NewUsers(nama string, umur int, alamat string) UserService {
	var user = entity.Users{
		Name:    nama,
		Age:     umur,
		Address: alamat,
	}

	return user
}

func NewCustomer(firstName string, lastName string, id int) entity.Customer {
	// var costomer = entity.Customer{
	// 	FirstName: firstName,
	// 	LastName:  lastName,
	// 	Id:        id,
	// }
	var costomer entity.Customer
	costomer.FirstName = firstName
	costomer.LastName = lastName
	costomer.Id = id

	return costomer
}
