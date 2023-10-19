package prototype

import "fmt"

type Shape interface {
	Print() string
	Clone() Shape
}

// ----------------------------------------------------------------
type Rectangle struct {
	name   string
	width  float64
	height float64
}

func (r *Rectangle) Print() string {
	return fmt.Sprintf("%s: %.2f %.2f", r.name, r.width, r.height)
}

func (r *Rectangle) Clone() Shape {
	newR := *r
	newR.name = newR.name + "_clone"
	return &newR
}

// ----------------------------------------------------------------
type Circle struct {
	name   string
	radius float64
}

func (c *Circle) Print() string {
	return fmt.Sprintf("%s: %.2f", c.name, c.radius)
}

func (c *Circle) Clone() Shape {
	newC := *c
	newC.name = newC.name + "_clone"
	return &newC
}

// ----------------------------------------------------------------
type ShapeList struct {
	name   string
	shapes []Shape
}

func (sl *ShapeList) Print() string {
	res := sl.name + "\n"

	for _, s := range sl.shapes {
		res += s.Print() + "\n"
	}

	return res
}

func (sl *ShapeList) Clone() Shape {
	var newSl ShapeList

	for _, s := range sl.shapes {
		newSl.shapes = append(newSl.shapes, s.Clone())
	}
	newSl.name = sl.name + "_clone"

	return &newSl
}
