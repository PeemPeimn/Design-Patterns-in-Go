package factory_method

type Dialog interface {
	Render() string
	CreateButton() Button
}

type windowsDialog struct {
	Dialog
}

func (d *windowsDialog) Render() string {
	button := d.CreateButton()
	return button.Render()
}

func (d *windowsDialog) CreateButton() Button {
	return &windowsButton{}
}

type webDialog struct {
	Dialog
}

func (d *webDialog) Render() string {
	button := d.CreateButton()
	return button.Render()
}

func (d *webDialog) CreateButton() Button {
	return &htmlButton{}
}
