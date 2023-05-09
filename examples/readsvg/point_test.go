package readsvg

import "testing"

func TestReadPoint(t *testing.T) {
	tables := []struct {
		data         string
		expected     Point
		dataExpected string
	}{
		// comma separated numbers
		{"3.0,-29", Point{3.0, -29.0}, ""},
		{"3.0,-29", Point{3.0, -29.0}, ""},
		{"  3.0,   -29", Point{3.0, -29.0}, ""},
		{"3.0,-29test", Point{3.0, -29.0}, "test"},
		{"3.0,-29 test", Point{3.0, -29.0}, " test"},
		{"3.22,-29", Point{3.22, -29.0}, ""},
		{"3.0   ,-29", Point{3.0, -29.0}, ""},
		{"-3.0   ,-29", Point{-3.0, -29.0}, ""},
		//space separated numbers
		{"3.0 -29", Point{3.0, -29.0}, ""},
		{"3.0 -29", Point{3.0, -29.0}, ""},
		{"  3.0   -29", Point{3.0, -29.0}, ""},
		{"3.0 -29test", Point{3.0, -29.0}, "test"},
		{"3.0 -29 test", Point{3.0, -29.0}, " test"},
		{"3.0 -29", Point{3.0, -29.0}, ""},
		{"-3.0   -29", Point{-3.0, -29.0}, ""},
	}

	for _, table := range tables {

		result, resultData, err := ReadPoint(table.data)

		if err != nil {
			t.Fatalf("reading from '%s' shouldn't return error '%s'", table.data, err.Error())
		}

		if result != table.expected {
			t.Errorf("From '%s' should be readed point '%s' but it's point '%s'", table.data, table.expected.String(), result.String())
		}

		if resultData != table.dataExpected {
			t.Errorf("After reading from '%s' should rest string '%s' but it's '%s'", table.data, table.dataExpected, resultData)
		}
	}
}

func TestReadPoint_errors(t *testing.T) {
	tables := []struct {
		data string
	}{
		{"d50"},
		{"50.00,h10"},
		{"50.00 h10"},
		{""},
	}

	for _, table := range tables {

		_, _, err := ReadPoint(table.data)

		if err == nil {
			t.Errorf("Error should be returned from reading point from '%s' string.", table.data)
		}

	}
}
