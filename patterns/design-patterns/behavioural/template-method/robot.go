package templatemethod

import "fmt"

type AutomotiveRobot struct{}

func (r *AutomotiveRobot) Test() {
	fmt.Println("Automotive test")
}

func (r *AutomotiveRobot) GetName() {
	fmt.Println("Automotive")
}

type CookieRobot struct{}

func (r *CookieRobot) Test() {
	fmt.Println("Cookie test")
}

func (r *CookieRobot) GetName() {
	fmt.Println("Cookie")
}
