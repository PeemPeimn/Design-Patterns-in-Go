package composite

type Component interface {
	GetPrice() int
}

type Product struct {
	price int
}

func (p *Product) GetPrice() int {
	return p.price
}

type Box struct {
	children []Component
}

func (b *Box) GetPrice() int {
	total := 0
	for _, c := range b.children {
		total += c.GetPrice()
	}
	return total
}
