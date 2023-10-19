package adapter

type Peg interface {
	GetRadius() float64
}

type SquarePeg struct {
	width float64
}

func (p *SquarePeg) GetWidth() float64 {
	return p.width
}
