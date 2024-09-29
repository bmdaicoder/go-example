package main

import (
	"fmt"
	"go-example/patterns/design-patterns/structural/flyweight"
)

func main() {
	factory := flyweight.NewTextStyleFactory()

	text1 := flyweight.NewText("Hello", factory.GetTextStyle("BOLD"))
	text2 := flyweight.NewText("World", factory.GetTextStyle("BOLD"))

	fmt.Println(text1.Format())
	fmt.Println(text2.Format())
}
