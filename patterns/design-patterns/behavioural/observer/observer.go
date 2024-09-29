package observer

import "fmt"

type Observer interface {
	OnNotify(Event)
}

type observer struct {
	id int
}

func NewObserver(id int) *observer {
	return &observer{id: id}
}

func (o *observer) OnNotify(e Event) {
	fmt.Printf("observer %d recieved event %d\n", o.id, e.Data)
}

type Event struct {
	Data int
}
