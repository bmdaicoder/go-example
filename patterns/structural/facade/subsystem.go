package facade

import "fmt"

type AudioPlayer struct{}

func (a *AudioPlayer) PlayAudio() {
	fmt.Println("Playing audio...")
}

type VideoPlayer struct{}

func (v *VideoPlayer) PlayVideo() {
	fmt.Println("Playing video...")
}

type ScreenManager struct{}

func (s *ScreenManager) ShowScreen() {
	fmt.Println("Showing screen...")
}
