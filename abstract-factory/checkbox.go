package abstract_factory

type CheckBox interface {
	Check() string
}

type windowsCheckBox struct{}

func (checkbox *windowsCheckBox) Check() string {
	return "WindowsCheckBox"
}

type macCheckBox struct{}

func (button *macCheckBox) Check() string {
	return "MacCheckBox"
}
