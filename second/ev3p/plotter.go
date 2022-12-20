package ev3p

import "fmt"

type Plotter struct {
	Width, Height int
	PenPosition   Point //Hold info about position of pen on paper
}

func NewPlotter() Plotter {
	p := new(Plotter)
	return *p
}

func (p *Plotter) PrintInfo() {
	fmt.Printf("Width : %d", p.Width)
	fmt.Printf("Height : %d", p.Height)
}
