package ev3p

import (
	"encoding/json"
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

func (p Point) String() string {
	return fmt.Sprintf("x=%v, y=%v", p.x, p.y)
}

func (point *Point) Eq(p *Point) bool {
	return p.x == point.x && p.y == point.y
}

func (point *Point) Subst(p *Point) Point {
	return Point{point.x - p.x, point.y - p.y}
}

func (point *Point) Distance(p *Point) int {
	t := point.Subst(p)
	return int(math.Round(math.Sqrt(float64(t.x*t.x) + float64(t.y*t.y))))
}

func (p *Point) GetX() int {
	return p.x
}

func (p *Point) GetY() int {
	return p.y
}

type privatePoint struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p Point) MarshalJSON() ([]byte, error) {
	return json.Marshal(privatePoint{p.x, p.y})
}

func (p *Point) UnmarshalJSON(b []byte) error {
	pp := new(privatePoint)
	json.Unmarshal(b, &pp)
	p.x = pp.X
	p.y = pp.Y
	return nil
}
