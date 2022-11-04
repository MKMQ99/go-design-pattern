package abstract_factory

import "fmt"

type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

type Winfactory struct{}

type Macfactory struct{}

func (Winfactory) CreateButton() Button {
	return &WinButton{}
}

func (Winfactory) CreateCheckbox() Checkbox {
	return &WinCheckbox{}
}

func (Macfactory) CreateButton() Button {
	return &MacButton{}
}

func (Macfactory) CreateCheckbox() Checkbox {
	return &MacCheckbox{}
}

type Button interface {
	paint() string
}

type WinButton struct{}

type MacButton struct{}

func (WinButton) paint() string {
	return fmt.Sprintf("%sstyle %s", "Win", "Button")
}

func (MacButton) paint() string {
	return fmt.Sprintf("%sstyle %s", "Mac", "Button")
}

type Checkbox interface {
	paint() string
}

type WinCheckbox struct{}

type MacCheckbox struct{}

func (WinCheckbox) paint() string {
	return fmt.Sprintf("%sstyle %s", "Win", "Checkbox")
}

func (MacCheckbox) paint() string {
	return fmt.Sprintf("%sstyle %s", "Mac", "Checkbox")
}
