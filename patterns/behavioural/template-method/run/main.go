package main

import templatemethod "go-example/patterns/behavioural/template-method"

func main() {
	automotiveRobot := &templatemethod.AutomotiveRobot{}
	robot := &templatemethod.Robot{
		Robot: automotiveRobot,
	}

	robot.Go()

	cookieRobot := &templatemethod.CookieRobot{}
	robot = &templatemethod.Robot{
		Robot: cookieRobot,
	}

	robot.Go()
}
