package adapter

type RoundHole struct {
	radius float64
}

func (rh *RoundHole) Fits(peg Peg) bool {
	return rh.radius >= peg.GetRadius()
}
