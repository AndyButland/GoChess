package main

import (
	"testing"
)

func TestIsKingInCheck(t *testing.T) {
	b := board{}
	b.init()

	var res bool
	res = b.isKingInCheck("B")
	if res {
		t.Errorf("King reported to be in check in initial position")
	}

	// Test 1: e4
	b.movePiece(square{file: "E", rank: 2}, square{file: "E", rank: 4})
	res = b.isKingInCheck("B")
	if res {
		t.Errorf("King reported to be in check in non-checking position")
	}

	// f6
	b.movePiece(square{file: "F", rank: 7}, square{file: "F", rank: 6})
	res = b.isKingInCheck("B")
	if res {
		t.Errorf("King reported to be in check in non-checking position")
	}

	// Qh5
	b.movePiece(square{file: "D", rank: 1}, square{file: "H", rank: 5})
	res = b.isKingInCheck("B")
	if !res {
		t.Errorf("King reported to be in not in check in checking position")
	}

	// g6
	b.movePiece(square{file: "G", rank: 7}, square{file: "G", rank: 6})
	res = b.isKingInCheck("B")
	if res {
		t.Errorf("King reported to be in check in non-checking position")
	}
}
