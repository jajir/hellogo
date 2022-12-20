package lev3

import log "github.com/sirupsen/logrus"

type Plotter struct {
	Width, Height int
	PenPosition   Point //Hold info about position of pen on paper
}

func NewPlotter() Plotter {
	p := new(Plotter)
	return *p
}

func (p *Plotter) PrintInfo() {
	log.Printf("Width : %d", p.Width)
	log.Printf("Height : %d", p.Height)
}
