package state

import (
	"fmt"
	"time"
)

type RedLight struct{}

func (l *RedLight) Enter() {
	fmt.Println("Red light is on. Stop driving")
	time.Sleep(time.Second * 3)
}

func (l *RedLight) Update(sm *StateMachine) {
	sm.setState(&GreenLight{})
}

type GreenLight struct{}

func (l *GreenLight) Enter() {
	fmt.Println("Green light is on. You can drive")
	time.Sleep(time.Second * 3)
}

func (l *GreenLight) Update(sm *StateMachine) {
	sm.setState(&YellowLight{})
}

type YellowLight struct{}

func (l *YellowLight) Enter() {
	fmt.Println("Yellow light is on. Prepare to stop")
	time.Sleep(time.Second * 3)
}

func (l *YellowLight) Update(sm *StateMachine) {
	sm.setState(&RedLight{})
}
