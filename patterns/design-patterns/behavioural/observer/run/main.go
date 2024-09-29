package main

import "go-example/patterns/design-patterns/behavioural/observer"

func main() {
	n := observer.NewNotifier()

	n.Register(observer.NewObserver(1))
	n.Register(observer.NewObserver(2))

	n.Notify(observer.Event{Data: 12})
	n.Notify(observer.Event{Data: 191})
	n.Notify(observer.Event{Data: 131})
	n.Notify(observer.Event{Data: 1256})
}
