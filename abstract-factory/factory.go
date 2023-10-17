package abstract_factory

type GUIFactory interface {
	CreateButton() Button
	CreateCheckBox() CheckBox
}

type windowsFactory struct{}

func (factory *windowsFactory) CreateButton() Button {
	return &windowsButton{}
}

func (factory *windowsFactory) CreateCheckBox() CheckBox {
	return &windowsCheckBox{}
}

type macFactory struct{}

func (factory *macFactory) CreateButton() Button {
	return &macButton{}
}

func (factory *macFactory) CreateCheckBox() CheckBox {
	return &macCheckBox{}
}
