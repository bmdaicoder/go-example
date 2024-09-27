package main

import (
	"fmt"
	"go-example/patterns/structural/composite"
)

func main() {
	file1 := &composite.File{Size: 200}
	file2 := &composite.File{Size: 300}

	dir := &composite.Directory{}
	dir.AddChild(file1)
	dir.AddChild(file2)

	fmt.Println(dir.GetSize())
}
