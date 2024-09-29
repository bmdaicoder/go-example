package observer

type Notifier interface {
	Register(Observer)
	UnRegister(Observer)
	Notifier(Event)
}

type notifier struct {
	observers map[Observer]struct{}
}

func NewNotifier() *notifier {
	return &notifier{
		observers: make(map[Observer]struct{}),
	}
}

func (n *notifier) Register(o Observer) {
	n.observers[o] = struct{}{}
}

func (n *notifier) Unregister(o Observer) {
	delete(n.observers, o)
}

func (n *notifier) Notify(e Event) {
	for o := range n.observers {
		o.OnNotify(e)
	}
}
