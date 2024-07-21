package areacalc

import "strings"

const pi = 3.14159

type Shape interface {
	Area() float64
	Type() string
}

type Rectangle struct {
	a float64
	b float64
	name string
}

func (rect Rectangle) Area() float64 {
	return rect.a * rect.b
}

func (rect Rectangle) Type() string {
	return rect.name
}

func NewRectangle(a float64, b float64, name string) *Rectangle {
	return &Rectangle{a, b, name}
}

type Circle struct {
	r float64
	name string
}

func (c Circle) Area() float64 {
	return pi * c.r * c.r
}

func (c Circle) Type() string {
	return c.name
}

func NewCircle(r float64, name string) *Circle {
	return &Circle{r, name}
}


func AreaCalculator(figures []Shape) (string, float64) {
	if len(figures) == 0 {
		return "", 0.0
	}

	ret := 0.0

	var sb strings.Builder

	ret += figures[0].Area()
	sb.WriteString(figures[0].Type())

	for _, fig := range figures[1:] {
		ret += fig.Area()
		sb.WriteString("-")
		sb.WriteString(fig.Type())
	}

	return sb.String(), ret
}
