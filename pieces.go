package main

import "fmt"

type piece interface {
	getName() string
	getLegalSquares(b board, sq square, color string) []square
}

type pawn struct{}
type rook struct{}
type knight struct{}
type bishop struct{}
type queen struct{}
type king struct{}

func (p pawn) getName() string   { return "P" }
func (p rook) getName() string   { return "R" }
func (p knight) getName() string { return "N" }
func (p bishop) getName() string { return "B" }
func (p queen) getName() string  { return "Q" }
func (p king) getName() string   { return "K" }

func (p pawn) getLegalSquares(b board, sq square, color string) []square {
	var squares []square
	var appended bool
	var direction int
	var secondRank int

	if color == "W" {
		direction = 1
		secondRank = 2
	} else {
		direction = -1
		secondRank = 7
	}

	// Single move forward - allowed if no blocking piece.
	if sq.rank >= 2 && sq.rank <= 7 {
		appended, squares = appendLegalSquare(squares, b, color, sq, 1*direction, false)
	}

	// Double move forward - allowed if single move was allowed, and on second rank.
	if appended && sq.rank == secondRank {
		appended, squares = appendLegalSquare(squares, b, color, sq, 2*direction, false)
	}

	return squares
}

func (p rook) getLegalSquares(b board, sq square, color string) []square {
	var res []square
	return res
}

func (p knight) getLegalSquares(b board, sq square, color string) []square {
	var res []square
	return res
}

func (p bishop) getLegalSquares(b board, sq square, color string) []square {
	var res []square
	return res
}

func (p queen) getLegalSquares(b board, sq square, color string) []square {
	var res []square
	return res
}

func (p king) getLegalSquares(b board, sq square, color string) []square {
	var res []square
	return res
}

func appendLegalSquare(squares []square, b board, color string, sq square, rankOffset int, canTakeOppositeColorPiece bool) (bool, []square) {
	newSquare := square{rank: sq.rank + rankOffset, file: sq.file}
	if b.isSquareEmpty(newSquare) {
		return true, append(squares, newSquare)
	}

	p, _ := b.getPieceAt(newSquare)
	if p.color != color && canTakeOppositeColorPiece {
		return true, append(squares, newSquare)
	}

	return false, squares
}

type coloredPiece struct {
	piece
	color string
}

func (cp coloredPiece) String() string {
	return fmt.Sprintf("%s%s", cp.color, cp.piece.getName())
}
