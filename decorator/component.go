package decorator

type Component interface {
	Execute() string
}

type ConcreteComponent struct{}

func (c *ConcreteComponent) Execute() string {
	return "execute."
}
