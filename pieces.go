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
		appended, _, squares = appendLegalSquare(squares, b, color, sq, 1*direction, 0, false)
	}

	// Double move forward - allowed if single move was allowed, and on second rank.
	if appended && sq.rank == secondRank {
		appended, _, squares = appendLegalSquare(squares, b, color, sq, 2*direction, 0, false)
	}

	return squares
}

func (p rook) getLegalSquares(b board, sq square, color string) []square {
	var squares []square
	var appended, willTakePiece bool

	// Vertically up from current position.
	for i := sq.rank + 1; i < BoardSize; i++ {
		appended, willTakePiece, squares = appendLegalSquare(squares, b, color, sq, i-sq.rank, 0, true)
		if !appended || (appended && willTakePiece) {
			break
		}
	}

	// Vertically down from current position.
	for i := sq.rank - 1; i > 0; i-- {
		appended, willTakePiece, squares = appendLegalSquare(squares, b, color, sq, i-sq.rank, 0, true)
		if !appended || (appended && willTakePiece) {
			break
		}
	}

	// Horizonally right from current position.
	fileNumber := fromFileStr(sq.file)
	for i := fileNumber + 1; i < BoardSize; i++ {
		appended, willTakePiece, squares = appendLegalSquare(squares, b, color, sq, 0, i-fileNumber, true)
		if !appended || (appended && willTakePiece) {
			break
		}
	}

	// Horizonally left from current position.
	for i := fileNumber - 1; i >= 0; i-- {
		appended, willTakePiece, squares = appendLegalSquare(squares, b, color, sq, 0, i-fileNumber, true)
		if !appended || (appended && willTakePiece) {
			break
		}
	}

	return squares
}

func (p knight) getLegalSquares(b board, sq square, color string) []square {
	var squares []square
	return squares
}

func (p bishop) getLegalSquares(b board, sq square, color string) []square {
	var squares []square
	return squares
}

func (p queen) getLegalSquares(b board, sq square, color string) []square {
	var squares []square
	return squares
}

func (p king) getLegalSquares(b board, sq square, color string) []square {
	var squares []square
	return squares
}

func appendLegalSquare(squares []square, b board, color string, sq square, rankOffset int, fileOffset int, canTakeOppositeColorPiece bool) (bool, bool, []square) {
	newSquare := square{rank: sq.rank + rankOffset, file: toFileStr(fromFileStr(sq.file) + fileOffset)}
	if b.isSquareEmpty(newSquare) {
		return true, false, append(squares, newSquare)
	}

	p, _ := b.getPieceAt(newSquare)
	if p.color != color && canTakeOppositeColorPiece {
		return true, true, append(squares, newSquare)
	}

	return false, false, squares
}

type coloredPiece struct {
	piece
	color string
}

func (cp coloredPiece) String() string {
	return fmt.Sprintf("%s%s", cp.color, cp.piece.getName())
}
