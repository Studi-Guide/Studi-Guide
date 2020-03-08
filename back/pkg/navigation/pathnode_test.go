package navigation

import (
	"studi-guide/ent"
	"testing"
)

func TestDistance(t *testing.T) {
	O := ent.PathNode{XCoordinate: 0, YCoordinate: 0, ZCoordinate: 0}
	p1 := ent.PathNode{XCoordinate: 1, YCoordinate: 2, ZCoordinate: 3}

	if distance := Distance(O, p1); distance != 4 {
		t.Error("expected: 4, got: ", distance)
	}

	if distance := Distance(p1, O); distance != 4 {
		t.Error("expected: 4, got: ", distance)
	}

	p2 := ent.PathNode{XCoordinate: 3, YCoordinate: 4, ZCoordinate: 5}
	if distance := Distance(p1, p2); distance != 3 {
		t.Error("expected: 3, got: ", distance)
	}

}