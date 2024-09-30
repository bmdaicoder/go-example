package mediator

import "fmt"

type Service interface {
	Process(event, metadata string)
}

type RegistrationService struct {
	eventMediator EventMediator
}

func NewRegistrationService(eventMediator EventMediator) RegistrationService {
	r := RegistrationService{eventMediator}
	eventMediator.RegisterService(r)
	return r
}

func (s RegistrationService) Process(event, metadata string) {
	if event == "activation-email-sent" {
		fmt.Println("RegistrationService -activation email sent to ", metadata)
	}
	if event == "activation-email-failed" {
		fmt.Println("RegistrationService - failed to send activation email to", metadata)
	}
}

func (s RegistrationService) RegisterUser(name string) {
	fmt.Println("RegistrationService - registered user", name)
	s.eventMediator.Fire("user-registered", name)
}

type EmailService struct {
	eventMediator EventMediator
}

func NewEmailService(eventMediator EventMediator) EmailService {
	r := EmailService{eventMediator}
	eventMediator.RegisterService(r)
	return r
}

func (s EmailService) Process(event string, metadata string) {
	if event == "user-registered" {
		if metadata == "anna" {
			fmt.Println("EmailService - sent activation email to", metadata)
			s.eventMediator.Fire("activation-email-sent", metadata)
		} else {
			fmt.Println("EmailService - failed to send to", metadata)
			s.eventMediator.Fire("activation-email-failed", metadata)
		}
	}
}
