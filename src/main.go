package main

import (
	"fmt"

	"go.mod/pkg/services"
)

func up(o int) interface{} {
	if o == 1 {
		return 1
	} else if o == 2 {
		return false
	} else {
		return "salah"
	}
}

func main() {

	// var datas interface{} = up(2)
	// fmt.Println(datas)
	user := services.NewUsers("anggi prabowo", 20, "Jakarta")
	fmt.Println(user.GetUsers())
	// costumer := services.NewCustomer("anggi", "prabowo", 18)
	// fmt.Println(costumer.GetFirstName())

}
