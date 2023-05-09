package readsvg

import (
	"fmt"
)

type Path struct {
	elements     []interface{}
	currentPoint Point
}

func (c Path) String() string {
	return fmt.Sprintf("Path[%s]", c.elements...)
}

type BesierQuadratic struct {
	P1, P2, P3 Point
}

func (c BesierQuadratic) String() string {
	return fmt.Sprintf("BesierQuadratic[%s, %s, %s]", c.P1.String(), c.P2.String(), c.P3.String())
}

func ParseD(data string) ([]interface{}, error) {
	path := Path{}
	data = TrimFromBeggining(data)
	if len(data) == 0 {
		return path.elements, nil
	}
	ch := data[0]
	data = TrimFromBeggining(data[1:])
	if ch == 'm' || ch == 'M' {
		currentPoint, data, _ := ReadPoint(data)
		path.currentPoint = currentPoint

		data = TrimFromBeggining(data)
		if len(data) == 0 {
			return nil, fmt.Errorf("can't parse path because it's empty '%s'", data)
		}
		command := data[0]

		if IsNumberOrSign(command) {
			//there is no command, jast coordinates, lets treat thema s lines
			data, err := ParseCommand('L', data, &path)
			if err != nil {
				return nil, fmt.Errorf("can't parse path '%s', because of %w", data, err)
			}
		} else {

			data = TrimFromBeggining(data[1:])
			if len(data) == 0 {
				return nil, fmt.Errorf("can't parse path because there are no data after first command '%s'", data)
			}

			data, err := ParseCommand(command, data, &path)
			if err != nil {
				return nil, fmt.Errorf("can't parse path '%s', because of %w", data, err)
			}

			err = readCommands(data, &path)
			if err != nil {
				return nil, fmt.Errorf("can't parse path '%s', because of %w", data, err)
			}
		}

	} else {
		return nil, fmt.Errorf("can't find first 'm' or 'M' in d attribute of path data '%s'", data)
	}
	return path.elements, nil
}

func readCommands(data string, path *Path) error {
	data = TrimFromBeggining(data)
	if len(data) == 0 {
		return nil
	}

	command := data[0]
	data = TrimFromBeggining(data[1:])
	if len(data) == 0 {
		return fmt.Errorf("can't parse path because there are no data after command '%c'", command)
	}

	data, err := ParseCommand(command, data, path)
	if err != nil {
		return fmt.Errorf("can't parse path '%s', because of %w", data, err)
	}

	return readCommands(data, path)
}

func ParseCommand(command byte, data string, path *Path) (string, error) {
	if command == 'C' || command == 'c' {
		data, err := ReadBesierCubic(data, path, command == 'c')
		if err != nil {
			return data, fmt.Errorf("can't parse Besiere Cubic from '%s', because of %w", data, err)
		}
		return data, nil
	} else if command == 'L' || command == 'l' {
		data, err := ReadLine(data, path, command == 'l')
		if err != nil {
			return data, fmt.Errorf("can't parse Besiere Cubic from '%s', because of %w", data, err)
		}
		return data, nil
	} else {
		return data, fmt.Errorf("unsupported command '%c' in data '%s'", command, data)
	}
}
