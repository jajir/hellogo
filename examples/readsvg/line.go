package readsvg

import (
	"errors"
	"fmt"
)

type Line struct {
	startPoint Point
	points     []Point
}

func NewLine(startPoint Point, points []Point) *Line {
	l := new(Line)
	l.startPoint = startPoint
	l.points = points
	return l
}

func (c *Line) String() string {
	return fmt.Sprintf("Line[%s, %v]", c.startPoint.String(), c.points)
}

func ReadLine(data string, path *Path, relative bool) (string, error) {
	data = TrimFromBeggining(data)
	if len(data) == 0 {
		return data, errors.New("can't read line from empty string")
	}

	line := Line{}
	line.startPoint = path.currentPoint
	ch := data[0]
	for IsNumberOrSign(ch) && len(data) > 0 {
		point, tmpData, err := ReadPoint(data)
		if err != nil {
			return tmpData, fmt.Errorf("can't read point for line object from '%s' string, because of %w", tmpData, err)
		}
		if relative {
			point.Add(path.currentPoint)
		}
		line.points = append(line.points, point)
		tmpData = TrimFromBeggining(tmpData)
		if len(tmpData) == 0 {
			ch = 'Q'
		} else {
			ch = tmpData[0]
		}
		data = tmpData
	}
	if len(line.points) == 0 {
		return data, fmt.Errorf("unable to parse line because there are no points in string '%s'", data)
	}
	path.currentPoint = line.points[len(line.points)-1]
	path.elements = append(path.elements, line)
	return data, nil
}
