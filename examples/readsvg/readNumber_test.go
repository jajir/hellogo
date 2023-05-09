package readsvg

import "testing"

func TestReadNumber(t *testing.T) {
	tables := []struct {
		data         string
		expected     float64
		dataExpected string
	}{
		{"50", 50.0, ""},
		{"50.", 50.0, ""},
		{"50.0", 50.0, ""},
		{"50.000", 50.0, ""},
		{"50.000", 50.0, ""},
		{"+50.000", 50.0, ""},
		{"-50.000", -50.0, ""},
		{"-50.000+", -50.0, "+"},
		{"-50.000-", -50.0, "-"},
		{"50.0 ", 50.0, " "},
		{"50.0  ", 50.0, "  "},
		{"50.0 76  ", 50.0, " 76  "},
		{"50.0A", 50.0, "A"},
		{" 50.0", 50.0, ""},
	}

	for _, table := range tables {

		result, resultData, _ := ReadNumber(table.data)

		if result != table.expected {
			t.Errorf("From string '%s' should be readed number '%f' but it's number '%f'", table.data, table.expected, result)
		}

		if resultData != table.dataExpected {
			t.Errorf("After reading from '%s' returned string should be '%s' but it's '%s'", table.data, table.dataExpected, resultData)
		}

	}
}

func TestReadNumber_errors(t *testing.T) {
	tables := []struct {
		data string
	}{
		{"d50"},
		{"50.00.10"},
		{""},
	}

	for _, table := range tables {

		_, _, err := ReadNumber(table.data)

		if err == nil {
			t.Errorf("Error should be returned from reading number from '%s' string.", table.data)
		}

	}
}

func TestTrimFromBeggining(t *testing.T) {
	tables := []struct {
		data     string
		expected string
	}{
		{"50", "50"},
		{" 50", "50"},
		{"    50", "50"},
		{"    50  ", "50  "},
		{"", ""},
		{"     ", ""},
	}

	for _, table := range tables {

		result := TrimFromBeggining(table.data)

		if result != table.expected {
			t.Errorf("After trimming '%s' string should be '%s' but it's '%s'", table.data, table.expected, result)
		}

	}
}
