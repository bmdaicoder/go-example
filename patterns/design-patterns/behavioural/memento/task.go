package memento

type List []string

type Tasks struct {
	list List
}

func NewTask() *Tasks {
	return &Tasks{make(List, 0)}
}

func (t *Tasks) Add(s string) {
	t.list = append(t.list, s)
}

func (t *Tasks) All() List {
	return t.list
}

func (t *Tasks) Memento() Memento {
	return Memento{list: t.list}
}

func (t *Tasks) Restore(m Memento) {
	t.list = m.list
}
