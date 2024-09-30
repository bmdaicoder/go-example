package mediator

import "container/list"

type EventMediator interface {
	Fire(event, metadata string)
	RegisterService(service Service)
}

type ListEventMediator struct {
	services *list.List
}

func NewListEventMediator() EventMediator {
	l := list.New()
	return ListEventMediator{l}
}

func (m ListEventMediator) Fire(event string, metadata string) {
	for element := m.services.Front(); element != nil; element = element.Next() {
		element.Value.(Service).Process(event, metadata)
	}
}

func (m ListEventMediator) RegisterService(service Service) {
	m.services.PushFront(service)
}
