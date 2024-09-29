package templatemethod

type RobotTemplate interface {
	Test()
	GetName()
}

type Robot struct {
	Robot RobotTemplate
}

func (r *Robot) Go() {
	r.Robot.Test()
	r.Robot.GetName()
}
