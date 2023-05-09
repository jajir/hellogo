package readsvg

import (
	"fmt"
)

type Point struct {
	X float64
	Y float64
}

func (c Point) String() string {
	return fmt.Sprintf("Point[%f, %f]", c.X, c.Y)
}

func (c *Point) Add(p Point) {
	c.X = c.X + p.X
	c.Y = c.Y + p.Y
}

func ReadPoint(data string) (Point, string, error) {
	// fmt.Println("read point " + data)
	X, data, err := ReadNumber(data)
	if err != nil {
		return Point{0, 0}, data, fmt.Errorf("can't read first part of point from '%s' string, because of %w", data, err)
	}
	ch := data[0]
	if ch == ',' {
		data = TrimFromBeggining(data[1:])
	} else if ch == ' ' {
		data = TrimFromBeggining(data[1:])
		ch := data[0]
		if ch == ',' {
			data = TrimFromBeggining(data[1:])
		}
	} else {
		return Point{0, 0}, data, fmt.Errorf("can't read second number of point from '%s' string", data)
	}

	Y, data, err := ReadNumber(data)
	if err != nil {
		return Point{0, 0}, data, fmt.Errorf("can't read second number of point from '%s' string, because of %w", data, err)
	}
	return Point{X, Y}, data, nil
}
