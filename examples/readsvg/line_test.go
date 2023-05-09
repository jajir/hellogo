package readsvg

import (
	"testing"
)

func TestReadLine(t *testing.T) {
	tables := []struct {
		data                 string
		isRelative           bool
		expectedCurrentPoint Point
		expected             []Point
		expectedData         string
	}{
		{" 345 178 ", false, Point{345, 178}, []Point{{345, 178}}, ""},
		{" 345 178 34 18", false, Point{34, 18}, []Point{{345, 178}, {34, 18}}, ""},
		{" 345,178 ", false, Point{345, 178}, []Point{{345, 178}}, ""},
		{" 345,178 34,18", false, Point{34, 18}, []Point{{345, 178}, {34, 18}}, ""},
		{" 345,178 34.22,18", false, Point{34.22, 18}, []Point{{345, 178}, {34.22, 18}}, ""},
		{" 345,178 -34.22,18", false, Point{-34.22, 18}, []Point{{345, 178}, {-34.22, 18}}, ""},
		{" 2,3 ", true, Point{3, 13}, []Point{{3, 13}}, ""},
		{" -2,3 ", true, Point{-1, 13}, []Point{{-1, 13}}, ""},
	}

	for _, table := range tables {
		currentPoint := Point{1, 10}
		path := Path{currentPoint: currentPoint}
		data, err := ReadLine(table.data, &path, table.isRelative)

		if err != nil {
			t.Fatalf("Parsing of '%s' shouldn't return error '%s'", table.data, err.Error())
		}

		if data != table.expectedData {
			t.Errorf("Parsing of '%s' should return data '%s' but it returns '%s' data",
				table.data, table.expectedData, data)
		}

		if table.expectedCurrentPoint != path.currentPoint {
			t.Errorf("Parsing of '%s' should return current point '%s' but it returns '%s' current point",
				table.data, path.currentPoint, table.expectedCurrentPoint)
		}

		if len(path.elements) != 1 {
			t.Fatalf("Parsing of '%s' should return path with 1 element but there is '%d' elements",
				table.data, len(path.elements))
		}
		line := (Line)(path.elements[0].(Line))

		if line.startPoint != currentPoint {
			t.Errorf("Parsing of '%s' should return start point '%s' but it returns '%s' start point",
				table.data, currentPoint, line.startPoint)
		}

		if len(table.expected) != len(line.points) {
			t.Fatalf("Parsing of '%s' should return '%d' point in line but it returns '%d' points",
				table.data, len(table.expected), len(line.points))
		}

		for i, expectedLine := range table.expected {
			returnedPoint := line.points[i]
			if returnedPoint != expectedLine {
				t.Errorf("Parsing of '%s' should at position '%d' return point '%s' but there is point '%s'",
					table.data, i, expectedLine, returnedPoint)
			}
		}

	}
}

func TestReadLine_errors(t *testing.T) {
	tables := []struct {
		data string
	}{
		{""},
	}

	for _, table := range tables {
		currentPoint := Point{1, 10}
		path := Path{currentPoint: currentPoint}
		_, err := ReadLine(table.data, &path, false)

		if err == nil {
			t.Errorf("Error should be returned from reading point from '%s' string.", table.data)
		}

	}
}
