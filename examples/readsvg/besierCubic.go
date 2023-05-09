package readsvg

import "fmt"

type BesierCubic struct {
	P1, P2, P3, P4 Point
}

func (c BesierCubic) String() string {
	return fmt.Sprintf("BesierCubic[%s, %s, %s, %s]", c.P1.String(), c.P2.String(), c.P3.String(), c.P4.String())
}

func ReadBesierCubic(data string, path *Path, relative bool) (string, error) {
	cubic := BesierCubic{P1: path.currentPoint}

	p2, data, err := ReadPoint(data)
	if err != nil {
		return data, fmt.Errorf("can't read second point of besiere cubic from '%s' string, because of %w",
			data, err)
	}
	cubic.P2 = p2

	p3, data, err := ReadPoint(data)
	if err != nil {
		return data, fmt.Errorf("can't read third point of besiere cubic from '%s' string, because of %w",
			data, err)
	}
	cubic.P3 = p3

	p4, data, err := ReadPoint(data)
	if err != nil {
		return data, fmt.Errorf("can't read forth point of besiere cubic from '%s' string, because of %w",
			data, err)
	}
	cubic.P4 = p4

	if relative {
		cubic.P2.Add(path.currentPoint)
		cubic.P3.Add(path.currentPoint)
		cubic.P4.Add(path.currentPoint)
	}

	path.elements = append(path.elements, cubic)
	path.currentPoint = cubic.P4
	return data, nil
}
