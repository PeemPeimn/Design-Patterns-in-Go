package adapter

import "math"

type PegAdapter struct {
	squarePeg *SquarePeg
}

func (p *PegAdapter) GetRadius() float64 {
	return p.squarePeg.GetWidth() / 2 * (math.Sqrt(2))
}
