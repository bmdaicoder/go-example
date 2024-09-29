package flyweight

import (
	"fmt"
)

type TextStyle struct {
	name string
}

func (t *TextStyle) Format(text string) string {
	return fmt.Sprintf("[%s %s]", t.name, text)
}

type TextStyleFactory struct {
	styles map[string]*TextStyle
}

func NewTextStyleFactory() *TextStyleFactory {
	return &TextStyleFactory{
		styles: make(map[string]*TextStyle),
	}
}

func (f *TextStyleFactory) GetTextStyle(name string) *TextStyle {
	if style, ok := f.styles[name]; ok {
		return style
	}

	style := &TextStyle{name: name}
	f.styles[name] = style

	return style
}

type Text struct {
	Text  string
	Style *TextStyle
}

func NewText(text string, style *TextStyle) *Text {
	return &Text{
		Text:  text,
		Style: style,
	}
}

func (t *Text) Format() string {
	return t.Style.Format(t.Text)
}
