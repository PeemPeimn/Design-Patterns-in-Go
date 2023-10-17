package abstract_factory

type Button interface {
	Paint() string
}

type windowsButton struct{}

func (button *windowsButton) Paint() string {
	return "WindowsButton"
}

type macButton struct{}

func (button *macButton) Paint() string {
	return "MacButton"
}
