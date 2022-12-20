package lev3

/*Picture is composed from connected points.
Picture could be stored to json and back and printed. */

import (
	"fmt"
)

type Picture struct {
	Min    Point //top left corner of picture
	Max    Point //buttom right corner of picture
	Points []Point
}

func (p Picture) String() string {
	return fmt.Sprintf("Picture{Min{%v}, Max{%v}, points=%v}", p.Min, p.Max, p.Points)
}

func (p Picture) GetStartPoint() Point {
	return p.Points[0]
}

func (p Picture) GetStepsCount() int {
	return len(p.Points) - 1
}

func (p Picture) GetStep(index int) Point {
	return p.Points[index+1]
}

func (p Picture) GetStepDiff(index int) Point {
	return p.Points[index+1].Subst(&p.Points[index])
}
