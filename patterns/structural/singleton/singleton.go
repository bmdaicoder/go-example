package singleton

import (
	"fmt"
	"sync"
)

type Singleton struct{}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	if instance == nil {
		once.Do(func() {
			instance = &Singleton{}
			fmt.Println("Creating instance.")
		})
	} else {
		fmt.Println("Instance already created.")
	}
	return instance
}
