package main

type Volumer interface {
	Volume() float64
}

type cube struct {
	edge float64
}

type triangularPrism struct {
	base     float64
	attitude float64
	height   float64
}

func (c cube) Volume() float64 {
	return c.edge * c.edge * c.edge
}

func (t triangularPrism) Volume() float64 {
	return 0.5 * t.attitude * t.height
}

func main() {

}
