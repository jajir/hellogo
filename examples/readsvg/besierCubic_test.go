package readsvg

import "testing"

func TestReadBesierCubic(t *testing.T) {
	tables := []struct {
		data         string
		isRelative   bool
		expected     BesierCubic
		expectedData string
	}{
		{"3,-29 77,-43 46,0", false, BesierCubic{Point{1, 10}, Point{3, -29}, Point{77, -43}, Point{46, 0}}, ""},
		{"3,-29 77,-43 46,6", true, BesierCubic{Point{1, 10}, Point{4, -19}, Point{78, -33}, Point{47, 16}}, ""},
	}

	for _, table := range tables {
		currentPoint := Point{1, 10}
		path := Path{currentPoint: currentPoint}
		data, err := ReadBesierCubic(table.data, &path, table.isRelative)

		if err != nil {
			t.Fatalf("Parsing of '%s' shouldn't return error '%s'", table.data, err.Error())
		}

		if data != table.expectedData {
			t.Errorf("Parsing of '%s' should left '%s' string but it left '%s' string", table.data, table.expectedData, data)
		}

		if len(path.elements) != 1 {
			t.Fatalf("Parsing of '%s' should return path with 1 element but there is '%d' elements",
				table.data, len(path.elements))
		}
		cubic := (BesierCubic)(path.elements[0].(BesierCubic))

		if table.expected.P1 != cubic.P1 {
			t.Errorf("Parsing of Cubic Besiere '%s' read first point '%s' but it should be '%s'",
				table.data, cubic.P1, table.expected.P1)
		}
		if table.expected.P2 != cubic.P2 {
			t.Errorf("Parsing of Cubic Besiere '%s' read second point '%s' but it should be '%s'",
				table.data, cubic.P2, table.expected.P2)
		}
		if table.expected.P3 != cubic.P3 {
			t.Errorf("Parsing of Cubic Besiere '%s' read third point '%s' but it should be '%s'",
				table.data, cubic.P3, table.expected.P3)
		}
		if table.expected.P4 != cubic.P4 {
			t.Errorf("Parsing of Cubic Besiere '%s' read fourth point '%s' but it should be '%s'",
				table.data, cubic.P4, table.expected.P4)
		}

	}
}

func TestReadBesierCubic_errors(t *testing.T) {
	tables := []struct {
		data string
	}{
		{"3,-29 77,-43"},
		{"3,-29"},
		{""},
	}

	for _, table := range tables {
		currentPoint := Point{1, 10}
		path := Path{currentPoint: currentPoint}
		_, err := ReadBesierCubic(table.data, &path, false)

		if err == nil {
			t.Errorf("Error should be returned from reading Quadratic Besiere curve from '%s' string.", table.data)
		}

	}
}
