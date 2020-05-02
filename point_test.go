package main

import (
	"github.com/jajir/hellogo/lev3"
	"testing"
)

func TestSubstract(t *testing.T) {
	tables := []struct {
		x1 int
		y1 int
		x2 int
		y2 int
		x3 int
		y3 int
	}{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 10, 0, -10},
		{0, 0, 10, 0, -10, 0},
		{100, 100, 200, 200, -100, -100},
	}

	for _, table := range tables {
		p1 := lev3.NewPoint(table.x1, table.y1)
		p2 := lev3.NewPoint(table.x2, table.y2)
		s := lev3.NewPoint(table.x3, table.y3)
		d := p1.Subst(&p2)
		if !s.Eq(&d) {
			t.Errorf("Substraction %v - %v should be %d but is %d", p1, p2, s, d)
		}
	}
}

func TestDistance(t *testing.T) {
	tables := []struct {
		x1 int
		y1 int
		x2 int
		y2 int
		d  int
	}{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 10, 10},
		{0, 0, 10, 0, 10},
		{100, 100, 200, 200, 141},
	}

	for _, table := range tables {
		p1 := lev3.NewPoint(table.x1, table.y1)
		p2 := lev3.NewPoint(table.x2, table.y2)
		d := p1.Distance(&p2)
		if d != table.d {
			t.Errorf("Distamce between %v and %v should be %d but is %d", p1, p2, table.d, d)
		}
	}
}
