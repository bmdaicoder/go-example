package main

import (
	"fmt"
	chainofreponsibility "go-example/patterns/design-patterns/behavioural/chain-of-reponsibility"
)

func main() {
	FirstReceiver := &chainofreponsibility.FirstReceiver{}
	SecondReceiver := &chainofreponsibility.SecondReceiver{}

	FirstReceiver.SetNext(SecondReceiver)

	request := "Hello"
	response := FirstReceiver.Handler(request)
	fmt.Println(response)
}
