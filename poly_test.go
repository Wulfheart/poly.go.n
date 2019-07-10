package polygo

import "testing"

func TestPolygonCreation(t *testing.T) {
	ps := []Point{{0, 0}, {0, 1}, {1, 1}, {1, 0}}
	poly := NewPolygon(ps...)
	for i, q := range ps {
		if poly.corners[i] != q {
			t.Error("Wrong coordinate taken.")
		}
	}
}

func TestPointInPolygon(t *testing.T) {
	type m struct {
		q     Point
		xpect bool
	}
	ps := []Point{{0, 0}, {0, 1}, {1, 1}, {1, 0}}
	poly := NewPolygon(ps...)
	tests := []m{
		{Point{0, 0}, true}, {Point{2, 8}, false}, {Point{0.5, 0.5}, true},
	}
	for _, j := range tests {
		if poly.Contains(j.q) != j.xpect {
			t.Error("Error on Point", j.q)
		}

	}
}
