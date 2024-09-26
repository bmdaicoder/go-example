package main

import "go-example/patterns/structural/adapter"

func main() {
	paypalGateway := adapter.NewPaypalAdapter()
	paypalGateway.ProcessPayment(15.5)

	amazonGateway := adapter.NewAmazonAdapter()
	amazonGateway.ProcessPayment(20.4)
}
