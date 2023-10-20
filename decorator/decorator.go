package decorator

type BaseDecorator struct {
	wrapee Component
}

func (b *BaseDecorator) Execute() string {
	return "execute_base." + b.wrapee.Execute()
}

type ConcreteDecorator1 struct {
	*BaseDecorator
}

func (d *ConcreteDecorator1) Execute() string {
	return "execute_dec1." + d.BaseDecorator.Execute()
}

type ConcreteDecorator2 struct {
	*ConcreteDecorator1
}

func (d *ConcreteDecorator2) Execute() string {
	return "execute_dec2." + d.ConcreteDecorator1.Execute()
}
