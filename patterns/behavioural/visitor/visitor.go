package visitor

type Visitor interface {
	VisitManager(manager *Manager)
	VisitEngineer(engineer *Engineer)
	VisitSalesman(salesman *Salesman)
}

type SalaryCalculator struct {
	TotalSalary float64
}

func (s *SalaryCalculator) VisitManager(manager *Manager) {
	s.TotalSalary += manager.Salary
}

func (s *SalaryCalculator) VisitEngineer(engineer *Engineer) {
	s.TotalSalary += engineer.Salary
}

func (s *SalaryCalculator) VisitSalesman(salesman *Salesman) {
	s.TotalSalary += salesman.Salary
}
