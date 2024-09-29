package main

import (
	"fmt"
	"go-example/patterns/design-patterns/behavioural/memento"
)

func main() {
	history := memento.NewHistory()
	tasks := memento.NewTask()

	tasks.Add("Task 1")
	history.Save(tasks.Memento())

	tasks.Add("Task 2")
	history.Save(tasks.Memento())

	fmt.Println(tasks.All())

	tasks.Restore(history.Undo())
	fmt.Println(tasks.All())

	tasks.Restore(history.Undo())
	fmt.Println(tasks.All())
}
