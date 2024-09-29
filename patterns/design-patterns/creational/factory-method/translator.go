package factorymethod

import "fmt"

type Translator interface {
	GetMessage() string
}

func GetTranslator(language string) (Translator, error) {
	switch language {
	case "english":
		return NewEnglishMessage(), nil
	case "french":
		return NewFrenchMessage(), nil
	default:
		return nil, fmt.Errorf("Unknown language")
	}
}

type EnglishMessage struct{}

func NewEnglishMessage() *EnglishMessage {
	return &EnglishMessage{}
}

func (m *EnglishMessage) GetMessage() string {
	return "Hello world"
}

type FrenchMessage struct{}

func NewFrenchMessage() *FrenchMessage {
	return &FrenchMessage{}
}

func (m *FrenchMessage) GetMessage() string {
	return "Bonjour le monde"
}
