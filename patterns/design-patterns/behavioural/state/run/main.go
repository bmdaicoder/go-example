package main

import "go-example/patterns/design-patterns/behavioural/state"

func main() {
	sm := state.NewStateMachine(&state.RedLight{})

	for {
		sm.Transition()
	}
}
