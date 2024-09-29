package main

import (
	"fmt"
	"go-example/patterns/design-patterns/structural/facade"
)

func main() {
	multimediaSystem := facade.NewMultimediaFacade()

	multimediaSystem.PlayMovie()
	fmt.Println("Playing a movie...")
}
