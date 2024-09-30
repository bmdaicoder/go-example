package main

import "go-example/patterns/design-patterns/behavioural/mediator"

func main() {
	eventMediator := mediator.NewListEventMediator()
	registrationService := mediator.NewRegistrationService(eventMediator)
	mediator.NewEmailService(eventMediator)

	registrationService.RegisterUser("anna")
	registrationService.RegisterUser("bob")
}
