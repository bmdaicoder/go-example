package main

import (
	"fmt"
	"go-example/patterns/design-patterns/creational/builder"
)

func main() {
	email := "bmdaicoder"
	balance := 127.5
	city := "HoChiMinh"

	responseBuilder := builder.NewResponseBuilder()
	responseBuilder.SetEmail(email)
	responseBuilder.SetBalance(balance)
	responseBuilder.SetCity(city)
	response := responseBuilder.Build()

	fmt.Println(response.Email)
	fmt.Println(response.Balance)
	fmt.Println(response.City)
}
