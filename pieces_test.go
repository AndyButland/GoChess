package main

import (
	"testing"
)

func TestPawngetLegalSquares(t *testing.T) {
	p := pawn{}
	b := board{}
	b.init()

	var sq square
	var res []square

	// Test: white pawn on second rank can move 1 or 2 squares
	sq = square{file: "E", rank: 2}
	res = p.getLegalSquares(b, sq, "W")
	if len(res) != 2 {
		t.Errorf("Expected white pawn on second rank to have 2 legal moves, but got: %d", len(res))
	}

	// Test: white pawn on third rank can move 1 square
	sq = square{file: "E", rank: 3}
	res = p.getLegalSquares(b, sq, "W")
	if len(res) != 1 {
		t.Errorf("Expected white pawn on third rank to have 1 legal moves, but got: %d", len(res))
	}

	// Test: black pawn on second rank can move 1 or 2 squares
	sq = square{file: "E", rank: 7}
	res = p.getLegalSquares(b, sq, "B")
	if len(res) != 2 {
		t.Errorf("Expected black pawn on second rank to have 2 legal moves, but got: %d", len(res))
	}

	// Test: black pawn on third rank can move 1 square
	sq = square{file: "E", rank: 6}
	res = p.getLegalSquares(b, sq, "B")
	if len(res) != 1 {
		t.Errorf("Expected black pawn on third rank to have 1 legal moves, but got: %d", len(res))
	}

	// Test: pawn can't move forward if there's a blocking piece
	b.movePiece(square{file: "E", rank: 7}, square{file: "E", rank: 3})
	sq = square{file: "E", rank: 2}
	res = p.getLegalSquares(b, sq, "W")
	if len(res) != 0 {
		t.Errorf("Expected white pawn on second rank with blocking piece to have 0 legal moves, but got: %d", len(res))
	}
	b.init()
}
