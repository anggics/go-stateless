package main

import (
	"fmt"

	"go.mod/pkg/services"
)

func main() {
	// user := servicess.NewUsers("anggi Prabowo", 18, "jakarta")
	// fmt.Println(user)
	costumer := services.NewCustomer("anggi", "prabowo", 18)
	fmt.Println(costumer)

}
