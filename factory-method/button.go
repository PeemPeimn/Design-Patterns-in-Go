package factory_method

type Button interface {
	Render() string
	OnClick()
}

type windowsButton struct {
	Button
}

func (b *windowsButton) Render() string {
  return "Windows button"
}

type htmlButton struct {
	Button
}

func (b *htmlButton) Render() string {
  return "HTML button"
}
