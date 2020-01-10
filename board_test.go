package main

import (
	"testing"
)

func TestAreSquaresAdjacent(t *testing.T) {
	var res bool
	res = areSquaresAdjacent(square{file: "A", rank: 2}, square{file: "A", rank: 3})
	if !res {
		t.Errorf("Squares reported as non-adjacent but are adjacent vertically.")
	}

	res = areSquaresAdjacent(square{file: "A", rank: 3}, square{file: "A", rank: 2})
	if !res {
		t.Errorf("Squares reported as non-adjacent but are adjacent vertically.")
	}

	res = areSquaresAdjacent(square{file: "A", rank: 2}, square{file: "B", rank: 2})
	if !res {
		t.Errorf("Squares reported as non-adjacent but are adjacent horizontally.")
	}

	res = areSquaresAdjacent(square{file: "B", rank: 2}, square{file: "A", rank: 2})
	if !res {
		t.Errorf("Squares reported as non-adjacent but are adjacent horizontally.")
	}

	res = areSquaresAdjacent(square{file: "A", rank: 2}, square{file: "B", rank: 3})
	if !res {
		t.Errorf("Squares reported as non-adjacent but are adjacent diagonally.")
	}

	res = areSquaresAdjacent(square{file: "B", rank: 3}, square{file: "A", rank: 2})
	if !res {
		t.Errorf("Squares reported as non-adjacent but are adjacent diagonally.")
	}

	res = areSquaresAdjacent(square{file: "A", rank: 2}, square{file: "A", rank: 4})
	if res {
		t.Errorf("Squares reported as adjacent but aren't.")
	}
}

func TestGetSquaresBetween(t *testing.T) {
	b := board{}

	var expectedCount int
	var res []square

	res = b.getSquaresBetween(square{file: "A", rank: 2}, square{file: "A", rank: 6})
	expectedCount = 3
	if len(res) != expectedCount {
		t.Errorf("Expected %d vertical squares between, but got: %d (%v)", expectedCount, len(res), res)
	}

	res = b.getSquaresBetween(square{file: "A", rank: 2}, square{file: "D", rank: 2})
	expectedCount = 2
	if len(res) != expectedCount {
		t.Errorf("Expected %d horizontal squares between, but got: %d (%v)", expectedCount, len(res), res)
	}

	res = b.getSquaresBetween(square{file: "A", rank: 1}, square{file: "H", rank: 8})
	expectedCount = 6
	if len(res) != expectedCount {
		t.Errorf("Expected %d diagonal squares between, but got: %d (%v)", expectedCount, len(res), res)
	}

	res = b.getSquaresBetween(square{file: "A", rank: 1}, square{file: "H", rank: 7})
	expectedCount = 0
	if len(res) != expectedCount {
		t.Errorf("Expected %d between for non aligned squares, but got: %d (%v)", expectedCount, len(res), res)
	}

	res = b.getSquaresBetween(square{file: "A", rank: 1}, square{file: "A", rank: 1})
	expectedCount = 0
	if len(res) != expectedCount {
		t.Errorf("Expected %d between for same square, but got: %d (%v)", expectedCount, len(res), res)
	}
}

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
	var reason string

	// Test: king not in check is not in check-mate
	res, reason = b.isKingInCheckMate("W")
	if res {
		t.Errorf("King reported to be in in check-mate when not in check. Reason: %s", reason)
	}

	// Test: king in check but can move
	b.movePiece(square{file: "F", rank: 2}, square{file: "F", rank: 3})
	b.movePiece(square{file: "E", rank: 7}, square{file: "E", rank: 6})
	b.movePiece(square{file: "D", rank: 2}, square{file: "D", rank: 3})
	b.movePiece(square{file: "D", rank: 8}, square{file: "H", rank: 4})
	res, reason = b.isKingInCheckMate("W")
	if res {
		t.Errorf("King reported to be in in check-mate when in check but could move. Reason: %s", reason)
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
	res, reason = b.isKingInCheckMate("W")
	if res {
		t.Errorf("King reported to be in in check-mate when in check and cannot move but checking piece can be taken. Reason: %s", reason)
	}
	b.init()

	// Test: king in check and cannot move, checking piece cannot be taken, may be able to block but
	// more than one checking piece
	b.movePiece(square{file: "F", rank: 2}, square{file: "F", rank: 3})
	b.movePiece(square{file: "E", rank: 7}, square{file: "E", rank: 6})

	b.movePiece(square{file: "E", rank: 2}, square{file: "E", rank: 3})
	b.movePiece(square{file: "G", rank: 8}, square{file: "F", rank: 6})

	b.movePiece(square{file: "E", rank: 1}, square{file: "F", rank: 2})
	b.movePiece(square{file: "F", rank: 6}, square{file: "H", rank: 5})

	b.movePiece(square{file: "D", rank: 1}, square{file: "E", rank: 2})
	b.movePiece(square{file: "H", rank: 5}, square{file: "G", rank: 3})

	b.movePiece(square{file: "A", rank: 2}, square{file: "A", rank: 3})
	b.movePiece(square{file: "D", rank: 8}, square{file: "H", rank: 4})

	b.movePiece(square{file: "B", rank: 2}, square{file: "B", rank: 3})
	b.movePiece(square{file: "G", rank: 3}, square{file: "E", rank: 4})
	res, reason = b.isKingInCheckMate("W")
	if !res {
		t.Errorf("King reported to not be in in check-mate but is, as more than one checking piece means check cannot be blocked. Reason: %s", reason)
	}
	b.init()

	// Test: king in check and cannot move, checking piece cannot be taken, may be able to block but
	// checking piece is a knight
	b.movePiece(square{file: "E", rank: 2}, square{file: "E", rank: 3})
	b.movePiece(square{file: "G", rank: 8}, square{file: "F", rank: 6})

	b.movePiece(square{file: "C", rank: 2}, square{file: "C", rank: 3})
	b.movePiece(square{file: "F", rank: 6}, square{file: "D", rank: 5})

	b.movePiece(square{file: "G", rank: 2}, square{file: "G", rank: 3})
	b.movePiece(square{file: "D", rank: 5}, square{file: "B", rank: 4})

	b.movePiece(square{file: "F", rank: 1}, square{file: "G", rank: 2})
	b.movePiece(square{file: "H", rank: 7}, square{file: "H", rank: 6})

	b.movePiece(square{file: "G", rank: 1}, square{file: "E", rank: 2})
	b.movePiece(square{file: "G", rank: 7}, square{file: "G", rank: 6})

	b.movePiece(square{file: "H", rank: 1}, square{file: "F", rank: 1})
	b.movePiece(square{file: "B", rank: 4}, square{file: "D", rank: 3})
	res, reason = b.isKingInCheckMate("W")
	if !res {
		t.Errorf("King reported to not be in in check-mate but is, as checking knight cannot be blocked. Reason: %s", reason)
	}
	b.init()

	// Test: king in check and cannot move, checking piece cannot be taken, can't block as checking
	// piece is adjacent.
	// Also tests that if king is only taker of checking piece, but king would be in check after
	// the taking, that this is still check-mate.
	b.movePiece(square{file: "E", rank: 2}, square{file: "E", rank: 3})
	b.movePiece(square{file: "E", rank: 7}, square{file: "E", rank: 5})

	b.movePiece(square{file: "E", rank: 1}, square{file: "E", rank: 2})
	b.movePiece(square{file: "E", rank: 5}, square{file: "E", rank: 4})

	b.movePiece(square{file: "D", rank: 1}, square{file: "E", rank: 1})
	b.movePiece(square{file: "F", rank: 7}, square{file: "F", rank: 5})

	b.movePiece(square{file: "D", rank: 2}, square{file: "D", rank: 3})
	b.movePiece(square{file: "F", rank: 5}, square{file: "F", rank: 4})

	b.movePiece(square{file: "C", rank: 1}, square{file: "D", rank: 2})
	b.movePiece(square{file: "A", rank: 7}, square{file: "A", rank: 6})

	b.movePiece(square{file: "B", rank: 1}, square{file: "C", rank: 3})
	b.movePiece(square{file: "B", rank: 7}, square{file: "B", rank: 6})

	b.movePiece(square{file: "A", rank: 1}, square{file: "D", rank: 1})
	b.movePiece(square{file: "C", rank: 7}, square{file: "C", rank: 6})

	b.movePiece(square{file: "G", rank: 2}, square{file: "G", rank: 3})
	b.movePiece(square{file: "D", rank: 7}, square{file: "D", rank: 6})

	b.movePiece(square{file: "G", rank: 1}, square{file: "H", rank: 3})
	b.movePiece(square{file: "F", rank: 4}, square{file: "F", rank: 3})
	res, reason = b.isKingInCheckMate("W")
	if !res {
		t.Errorf("King reported to not be in in check-mate but is, as checking piece is adjacent and can't be blocked. Reason: %s", reason)
	}
	b.init()

	// // Test: king in check and cannot move, checking piece cannot be taken, can block
	// res, reason = b.isKingInCheckMate("W")
	// if res {
	// 	t.Errorf("King reported to be in in check-mate when in check and cannot move but checking piece can be blocked. Reason: %s", reason)
	// }
	// b.init()

	// // Test: king in check and cannot move, checking piece cannot be taken, can block
	// res, reason = b.isKingInCheckMate("W")
	// if !res {
	// 	t.Errorf("King reported to not be in in check-mate but is. Reason: %s", reason)
	// }
}
