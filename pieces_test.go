package main

import (
	"testing"
)

func TestPawnGetLegalSquares(t *testing.T) {
	p := pawn{}
	b := board{}
	b.init()

	var sq square
	var res []square
	var expectedCount int

	// Test: white pawn on second rank can move 1 or 2 squares
	sq = square{file: "E", rank: 2}
	res = p.getLegalSquares(b, sq, "W")
	expectedCount = 2
	if len(res) != expectedCount {
		t.Errorf("Expected white pawn on second rank to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}

	// Test: white pawn on third rank can move 1 square
	sq = square{file: "E", rank: 3}
	res = p.getLegalSquares(b, sq, "W")
	expectedCount = 1
	if len(res) != expectedCount {
		t.Errorf("Expected white pawn on third rank to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}

	// Test: black pawn on second rank can move 1 or 2 squares
	sq = square{file: "E", rank: 7}
	res = p.getLegalSquares(b, sq, "B")
	expectedCount = 2
	if len(res) != expectedCount {
		t.Errorf("Expected black pawn on second rank to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}

	// Test: black pawn on third rank can move 1 square
	sq = square{file: "E", rank: 6}
	res = p.getLegalSquares(b, sq, "B")
	expectedCount = 1
	if len(res) != expectedCount {
		t.Errorf("Expected black pawn on third rank to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}

	// Test: pawn can't move forward if there's a blocking piece
	b.movePiece(square{file: "E", rank: 7}, square{file: "E", rank: 3})
	sq = square{file: "E", rank: 2}
	res = p.getLegalSquares(b, sq, "W")
	expectedCount = 0
	if len(res) != expectedCount {
		t.Errorf("Expected white pawn on second rank with blocking piece to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}
	b.init()
}

func TestRookGetLegalSquares(t *testing.T) {
	p := rook{}
	b := board{}
	b.init()

	var sq square
	var res []square
	var expectedCount int

	// Test: white rook in starting position has no legal squares
	sq = square{file: "A", rank: 1}
	res = p.getLegalSquares(b, sq, "W")
	expectedCount = 0
	if len(res) != expectedCount {
		t.Errorf("Expected white rook in starting position to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}

	// Test: white rook with spaces around (on B4 of otherwise initialised board) has legal moves:
	// - 3 vertical above (two empty, one take of opponent pawn)
	// - 1 vertical below (empty)
	// - 7 horizontally (all empty)
	b.movePiece(square{file: "A", rank: 1}, square{file: "B", rank: 4})
	sq = square{file: "B", rank: 4}
	res = p.getLegalSquares(b, sq, "W")
	expectedCount = 11
	if len(res) != expectedCount {
		t.Errorf("Expected white rook with spaces around to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}
	b.init()

	// Test: black rook in initial position with spaces in front due to pawn move around has legal moves:
	// - 2 vertical above (two empty, vacated by pawn)
	b.movePiece(square{file: "H", rank: 7}, square{file: "H", rank: 5})
	sq = square{file: "H", rank: 8}
	res = p.getLegalSquares(b, sq, "B")
	expectedCount = 2
	if len(res) != expectedCount {
		t.Errorf("Expected black rook in initial position with spaces in front due to pawn move to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}
	b.init()
}

func TestKnightGetLegalSquares(t *testing.T) {
	p := knight{}
	b := board{}
	b.init()

	var sq square
	var res []square
	var expectedCount int

	// Test: white knight in starting position has 2 legal squares
	sq = square{file: "B", rank: 1}
	res = p.getLegalSquares(b, sq, "W")
	expectedCount = 2
	if len(res) != expectedCount {
		t.Errorf("Expected white knight in starting position to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}

	// Test: white knight with spaces around (on B4 of otherwise initialised board) has legal moves:
	b.movePiece(square{file: "B", rank: 1}, square{file: "B", rank: 4})
	sq = square{file: "B", rank: 4}
	res = p.getLegalSquares(b, sq, "W")
	expectedCount = 4
	if len(res) != expectedCount {
		t.Errorf("Expected white knight with spaces around to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}
	b.init()

	// Test: white knight with spaces around (on B5 of otherwise initialised board) has legal moves:
	b.movePiece(square{file: "B", rank: 1}, square{file: "B", rank: 5})
	sq = square{file: "B", rank: 5}
	res = p.getLegalSquares(b, sq, "W")
	expectedCount = 6
	if len(res) != expectedCount {
		t.Errorf("Expected white knight with spaces around to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}
	b.init()
}

func TestBishopGetLegalSquares(t *testing.T) {
	p := bishop{}
	b := board{}
	b.init()

	var sq square
	var res []square
	var expectedCount int

	// Test: white bishop in starting position has no legal squares
	sq = square{file: "C", rank: 1}
	res = p.getLegalSquares(b, sq, "W")
	expectedCount = 0
	if len(res) != expectedCount {
		t.Errorf("Expected white bishop in starting position to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}

	// Test: white bishop with spaces around (on B4 of otherwise initialised board) has legal moves:
	// - 3 up/right (two empty, one take of opponent pawn)
	// - 1 up/left (empty)
	// - 1 down/right (empty)
	// - 1 down/left (empty)
	b.movePiece(square{file: "C", rank: 1}, square{file: "B", rank: 4})
	sq = square{file: "B", rank: 4}
	res = p.getLegalSquares(b, sq, "W")
	expectedCount = 6
	if len(res) != expectedCount {
		t.Errorf("Expected white bishop with spaces around to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}
	b.init()

	// Test: white bishop in initial position with spaces available due to king pawn move around has legal moves:
	// - 5 up/left
	b.movePiece(square{file: "E", rank: 2}, square{file: "E", rank: 4})
	sq = square{file: "F", rank: 1}
	res = p.getLegalSquares(b, sq, "W")
	expectedCount = 5
	if len(res) != expectedCount {
		t.Errorf("Expected white bishop in initial position with spaces available due to queen pawn move to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}
	b.init()

	// Test: black bishop in initial position with spaces available due to queen pawn move around has legal moves:
	// - 5 down/right
	b.movePiece(square{file: "D", rank: 7}, square{file: "D", rank: 5})
	sq = square{file: "C", rank: 8}
	res = p.getLegalSquares(b, sq, "B")
	expectedCount = 5
	if len(res) != expectedCount {
		t.Errorf("Expected black bishop in initial position with spaces available due to queen pawn move to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}
	b.init()
}

func TestQueenGetLegalSquares(t *testing.T) {
	p := queen{}
	b := board{}
	b.init()

	var sq square
	var res []square
	var expectedCount int

	// Test: white queen in starting position has no legal squares
	sq = square{file: "D", rank: 1}
	res = p.getLegalSquares(b, sq, "W")
	expectedCount = 0
	if len(res) != expectedCount {
		t.Errorf("Expected white queen in starting position to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}

	// Test: white queen with spaces around (on B4 of otherwise initialised board) has legal moves:
	// - 3 vertical above (two empty, one take of opponent pawn)
	// - 1 vertical below (empty)
	// - 7 horizontally (all empty)
	// - 3 up/right (two empty, one take of opponent pawn)
	// - 1 up/left (empty)
	// - 1 down/right (empty)
	// - 1 down/left (empty)
	b.movePiece(square{file: "D", rank: 1}, square{file: "B", rank: 4})
	sq = square{file: "B", rank: 4}
	res = p.getLegalSquares(b, sq, "W")
	expectedCount = 17
	if len(res) != expectedCount {
		t.Errorf("Expected white queen with spaces around to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}
	b.init()
}

func TestKingGetLegalSquares(t *testing.T) {
	p := king{}
	b := board{}
	b.init()

	var sq square
	var res []square
	var expectedCount int

	// Test: white king in starting position has no legal squares
	sq = square{file: "E", rank: 1}
	res = p.getLegalSquares(b, sq, "W")
	expectedCount = 0
	if len(res) != expectedCount {
		t.Errorf("Expected white king in starting position to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}

	// Test: white king with spaces (on B4 of otherwise initialised board) around has legal moves:
	// - 8 empty squares around
	b.movePiece(square{file: "E", rank: 1}, square{file: "B", rank: 4})
	sq = square{file: "B", rank: 4}
	res = p.getLegalSquares(b, sq, "W")
	expectedCount = 8
	if len(res) != expectedCount {
		t.Errorf("Expected white king with spaces around to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}
	b.init()

	// Test: white king with spaces (on B3 of otherwise initialised board) around has legal moves:
	// - 5 empty squares around
	// - 3 blocked by own pawns
	b.movePiece(square{file: "E", rank: 1}, square{file: "B", rank: 3})
	sq = square{file: "B", rank: 3}
	res = p.getLegalSquares(b, sq, "W")
	expectedCount = 5
	if len(res) != expectedCount {
		t.Errorf("Expected white king with spaces around to have %d legal moves, but got: %d (%v)", expectedCount, len(res), res)
	}
	b.init()
}
