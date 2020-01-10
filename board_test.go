package main

import (
	"testing"
)

func TestIsKingInCheck(t *testing.T) {
	b := board{}
	b.init()

	var res bool
	var checkingSquares []square
	res, checkingSquares = b.isKingInCheck("B")
	if res {
		t.Errorf("King reported to be in check in initial position")
	}

	// Test 1: e4
	b.movePiece(square{file: "E", rank: 2}, square{file: "E", rank: 4})
	res, checkingSquares = b.isKingInCheck("B")
	if res {
		t.Errorf("King reported to be in check in non-checking position")
	}

	// f6
	b.movePiece(square{file: "F", rank: 7}, square{file: "F", rank: 6})
	res, checkingSquares = b.isKingInCheck("B")
	if res {
		t.Errorf("King reported to be in check in non-checking position")
	}

	// Qh5
	b.movePiece(square{file: "D", rank: 1}, square{file: "H", rank: 5})
	res, checkingSquares = b.isKingInCheck("B")
	if !res {
		t.Errorf("King reported to be in not in check in checking position")
	}
	if len(checkingSquares) != 1 {
		t.Errorf("King in check but no checking pieces found")
	}

	// g6
	b.movePiece(square{file: "G", rank: 7}, square{file: "G", rank: 6})
	res, checkingSquares = b.isKingInCheck("B")
	if res {
		t.Errorf("King reported to be in check in non-checking position")
	}
}

func TestIsKingInCheckMate(t *testing.T) {
	b := board{}
	b.init()

	var res bool

	// Test: king not in check is not in check-mate
	res = b.isKingInCheckMate("W")
	if res {
		t.Errorf("King reported to be in in check-mate when not in check")
	}

	// Test: king in check but can move
	b.movePiece(square{file: "F", rank: 2}, square{file: "F", rank: 3})
	b.movePiece(square{file: "E", rank: 7}, square{file: "E", rank: 6})
	b.movePiece(square{file: "D", rank: 2}, square{file: "D", rank: 3})
	b.movePiece(square{file: "D", rank: 8}, square{file: "H", rank: 4})
	res = b.isKingInCheckMate("W")
	if res {
		t.Errorf("King reported to be in in check-mate when in check but could move")
	}
	b.init()

	// Test: king in check and cannot move, checking piece can be taken
	b.movePiece(square{file: "F", rank: 2}, square{file: "F", rank: 3})
	b.movePiece(square{file: "E", rank: 7}, square{file: "E", rank: 6})
	b.movePiece(square{file: "A", rank: 2}, square{file: "A", rank: 4})
	b.movePiece(square{file: "A", rank: 7}, square{file: "A", rank: 6})
	b.movePiece(square{file: "A", rank: 4}, square{file: "A", rank: 5})
	b.movePiece(square{file: "B", rank: 7}, square{file: "B", rank: 6})
	b.movePiece(square{file: "A", rank: 1}, square{file: "A", rank: 4})
	b.movePiece(square{file: "D", rank: 8}, square{file: "H", rank: 4})
	res = b.isKingInCheckMate("W")
	if res {
		t.Errorf("King reported to be in in check-mate when in check and cannot move but checing piece can be taken")
	}

	// Test: king in check and cannot move, checking piece cannot be taken, but can block
	// TODO:
}
