package visitor

type Employee interface {
	Accept(visitor Visitor)
}

type Manager struct {
	Name   string
	Salary float64
}

func (m *Manager) Accept(visitor Visitor) {
	visitor.VisitManager(m)
}

type Engineer struct {
	Name   string
	Salary float64
}

func (e *Engineer) Accept(visitor Visitor) {
	visitor.VisitEngineer(e)
}

type Salesman struct {
	Name   string
	Salary float64
}

func (s *Salesman) Accept(visitor Visitor) {
	visitor.VisitSalesman(s)
}
