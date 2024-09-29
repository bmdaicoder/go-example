package main

import (
	"fmt"
	"go-example/patterns/behavioural/strategy"
)

func main() {
	array := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}

	bubbleSort := &strategy.BubbleSort{}
	insertionSort := &strategy.InsertionSort{}

	context := &strategy.Context{}
	context.SetSorter(bubbleSort)
	fmt.Println("Bubble Sort:", context.ExecuteSort(array))

	context.SetSorter(insertionSort)
	fmt.Println("Insertion Sort:", context.ExecuteSort(array))
}
