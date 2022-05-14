package lev3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
)

type Lissajous struct {
	maxX    float64 //max width of image
	maxY    float64 //max height of image
	cyclesX float64 // how many peeks should curve have in in X axe
	cyclesY float64 // how many peeks should curve have in in Y axe
	points  []Point
}

func NewLissajous(maxX, maxY, cyclesX, cyclesY int) Lissajous {
	var out Lissajous = *new(Lissajous)
	out.maxX = float64(maxX)
	out.maxY = float64(maxY)
	out.cyclesX = float64(cyclesX)
	out.cyclesY = float64(cyclesY)
	out.points = make([]Point, 0, 32000)
	out.Prepare()
	return out
}

func (l *Lissajous) PrintInfo() {
	fmt.Printf("len=%d cap=%d\n", len(l.points), cap(l.points))
	//	fmt.Printf("len=%d cap=%d %v\n", len(l.points), cap(l.points), l.points)
}

func (l *Lissajous) Prepare() {
	step := 0.0001
	halfX := l.maxX / 2
	halfY := l.maxY / 2
	for c := 0.0; c < math.Pi*l.cyclesX*l.cyclesY+1; c += step {
		x := math.Cos(l.cyclesX*c)*halfX + halfX
		y := math.Sin(l.cyclesY*c)*halfY + halfY
		point := NewPoint(int(x), int(y))
		if len(l.points) > 0 {
			last := l.points[len(l.points)-1]
			if point.Distance(&last) > 5 {
				l.points = append(l.points, point)
			}
		} else {
			l.points = append(l.points, point)
		}
	}
}

func (l Lissajous) GetStartPoint() Point {
	return l.points[0]
}

func (l Lissajous) GetStepsCount() int {
	return len(l.points) - 1
}

func (l Lissajous) GetStep(index int) Point {
	return l.points[index+1]
}

func (l Lissajous) GetStepDiff(index int) Point {
	return l.points[index+1].Subst(&l.points[index])
}

func (l *Lissajous) Marshal() (io.Reader, error) {
	b, err := json.MarshalIndent(l, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (l *Lissajous) GetPicture() Picture {
	var out Picture = *new(Picture)
	out.Min = Point{0, 0}
	out.Max = Point{int(l.maxX), int(l.maxY)}
	out.Points = make([]Point, len(l.points))
	copy(out.Points, l.points)
	return out
}
