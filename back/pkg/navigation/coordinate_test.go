package navigation

import "testing"

func TestDistanceTo(t *testing.T) {
	O := Coordinate{X: 0, Y: 0, Z: 0}
	p1 := Coordinate{X: 1, Y: 2, Z: 3}

	if distance := O.DistanceTo(p1); distance != 4 {
		t.Error("expected: 4, got: ", distance)
	}

	if distance := p1.DistanceTo(O); distance != 4 {
		t.Error("expected: 4, got: ", distance)
	}

	p2 := Coordinate{X: 3, Y: 4, Z: 5}
	if distance := p2.DistanceTo(p1); distance != 3 {
		t.Error("expected: 3, got: ", distance)
	}

}

func TestEquals(t *testing.T) {
	p1 := Coordinate{X: 1, Y: 2, Z: 3}
	p2 := Coordinate{X: 3, Y: 2, Z: 3}
	p3 := Coordinate{X: 1, Y: 2, Z: 3}

	if !p1.Equals(p1) {
		t.Error("expected p1 equals self")
	}

	if p2.Equals(p1) {
		t.Error("expected p2 not equals p1")
	}

	if !p3.Equals(p1) {
		t.Error("expected p3 equals p2")
	}
}