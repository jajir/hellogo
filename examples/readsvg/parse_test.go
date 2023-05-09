package readsvg

import (
	"reflect"
	"testing"
)

func TestParseD(t *testing.T) {
	tables := []struct {
		data     string
		expected []interface{}
	}{
		{"", []interface{}{}},
		{"M1 3 L4 5 ", []interface{}{Line{Point{1, 3}, []Point{{4, 5}}}}},
		{"M1 3 l4 5 ", []interface{}{Line{Point{1, 3}, []Point{{5, 8}}}}},
		{"m 686.622,962.563 -1.234,0.296 -0.982,0.822 -0.663,1.209 -0.004,1.663 ", []interface{}{
			Line{Point{686.622, 962.563}, []Point{{-1.234, 0.296}, {-0.982, 0.822}, {-0.663, 1.209}, {-0.004, 1.663}}}}},
	}

	for _, table := range tables {

		result, err := ParseD(table.data)

		if err != nil {
			t.Fatalf("Parsing of '%s' shouldn't return error '%s'", table.data, err.Error())
		}

		if len(result) != len(table.expected) {
			t.Fatalf("Parsing of '%s' should return '%d' objects but it returns '%d' objects", table.data, len(table.expected), len(result))
		}

		for i, expectedObject := range table.expected {
			returnedObject := result[i]

			returnedType := reflect.TypeOf(returnedObject)
			expectedType := reflect.TypeOf(expectedObject)
			if returnedType != expectedType {
				t.Fatalf("Parsing of '%s' should at position '%d' returned object which type is '%s' but returned is '%s'",
					table.data, i, expectedType, returnedType)
			}
			if !reflect.DeepEqual(returnedObject, expectedObject) {
				t.Errorf("Parsing of '%s' should at position '%d' return object '%s' but it's object '%s'",
					table.data, i, expectedObject, returnedObject)
			}
		}

	}
}

func TestParseD_errors(t *testing.T) {
	tables := []struct {
		data string
	}{
		{"d50"},
		{"50.00,h10"},
		{"50.00 h10"},
		{"M3,5 "},
		{"M3,5 C"},
	}

	for _, table := range tables {

		_, err := ParseD(table.data)

		if err == nil {
			t.Errorf("Error should be returned from reading point from '%s' string.", table.data)
		}

	}
}
