package main

import "go-example/patterns/design-patterns/behavioural/command"

func main() {
	tv := &command.Tv{}

	onCommand := &command.OnCommand{
		Device: tv,
	}

	offCommand := &command.OffCommand{
		Device: tv,
	}

	onButton := &command.Button{
		Command: onCommand,
	}
	onButton.Press()

	offButton := &command.Button{
		Command: offCommand,
	}
	offButton.Press()
}
