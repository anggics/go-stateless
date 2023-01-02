package services

import "go.mod/pkg/entity"

func NewCustomer(firstName string, lastName string, id int) entity.Customer {
	var costomer entity.Customer
	costomer.FirstName = firstName
	costomer.LastName = lastName
	costomer.Id = id

	return costomer
}
