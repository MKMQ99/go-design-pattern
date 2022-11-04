package abstract_factory

import "testing"

func TestAbstractFactory(t *testing.T) {
	var factory GUIFactory
	factory = Winfactory{}
	s := factory.CreateButton().paint()
	if s != "Winstyle Button" {
		t.Fatal("Winstyle Button Error")
	}
	s = factory.CreateCheckbox().paint()
	if s != "Winstyle Checkbox" {
		t.Fatal("Winstyle Checkbox Error")
	}
	factory = Macfactory{}
	s = factory.CreateButton().paint()
	if s != "Macstyle Button" {
		t.Fatal("Macstyle Button Error")
	}
	s = factory.CreateCheckbox().paint()
	if s != "Macstyle Checkbox" {
		t.Fatal("Macstyle Checkbox Error")
	}
}
