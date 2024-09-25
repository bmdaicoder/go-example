package main

import (
	"fmt"
	factorymethod "go-example/patterns/creational/factory-method"
)

func main() {
	translator, err := factorymethod.GetTranslator("english")
	if err != nil {
		return
	}

	fmt.Println(translator.GetMessage())

	translator, err = factorymethod.GetTranslator("french")
	if err != nil {
		return
	}

	fmt.Println(translator.GetMessage())
}
