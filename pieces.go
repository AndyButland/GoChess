package main

import "fmt"

type takingBehavior int

const (
	CannotTake takingBehavior = iota
	CanTake
	MustTake
)

type piece interface {
	getName() string
	getLegalSquares(b board, sq square, color string) []square
}

type coloredPiece struct {
	piece
	color string
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
		appended, _, squares = appendLegalSquare(squares, b, color, sq, 1*direction, 0, CannotTake)
	}

	// Double move forward - allowed if single move was allowed, and on second rank.
	if appended && sq.rank == secondRank {
		appended, _, squares = appendLegalSquare(squares, b, color, sq, 2*direction, 0, CannotTake)
	}

	// Diagonal taking moves
	_, _, squares = appendLegalSquare(squares, b, color, sq, 1*direction, 1, MustTake)
	_, _, squares = appendLegalSquare(squares, b, color, sq, 1*direction, -1, MustTake)

	return squares
}

func (p rook) getLegalSquares(b board, sq square, color string) []square {
	return getLegalSquaresForRook(b, sq, color)
}

func getLegalSquaresForRook(b board, sq square, color string) []square {
	var squares []square
	var appended, willTakePiece bool

	// Vertically up from current position.
	for i := sq.rank + 1; i < BoardSize; i++ {
		appended, willTakePiece, squares = appendLegalSquare(squares, b, color, sq, i-sq.rank, 0, CanTake)
		if !appended || (appended && willTakePiece) {
			break
		}
	}

	// Vertically down from current position.
	for i := sq.rank - 1; i > 0; i-- {
		appended, willTakePiece, squares = appendLegalSquare(squares, b, color, sq, i-sq.rank, 0, CanTake)
		if !appended || (appended && willTakePiece) {
			break
		}
	}

	// Horizonally right from current position.
	fileNumber := fromFileStr(sq.file)
	for i := fileNumber + 1; i < BoardSize; i++ {
		appended, willTakePiece, squares = appendLegalSquare(squares, b, color, sq, 0, i-fileNumber, CanTake)
		if !appended || (appended && willTakePiece) {
			break
		}
	}

	// Horizonally left from current position.
	for i := fileNumber - 1; i >= 0; i-- {
		appended, willTakePiece, squares = appendLegalSquare(squares, b, color, sq, 0, i-fileNumber, CanTake)
		if !appended || (appended && willTakePiece) {
			break
		}
	}

	return squares
}

func (p knight) getLegalSquares(b board, sq square, color string) []square {
	var squares []square

	_, _, squares = appendLegalSquare(squares, b, color, sq, 2, 1, CanTake)
	_, _, squares = appendLegalSquare(squares, b, color, sq, 1, 2, CanTake)
	_, _, squares = appendLegalSquare(squares, b, color, sq, -1, 2, CanTake)
	_, _, squares = appendLegalSquare(squares, b, color, sq, -2, 1, CanTake)
	_, _, squares = appendLegalSquare(squares, b, color, sq, -2, -1, CanTake)
	_, _, squares = appendLegalSquare(squares, b, color, sq, -1, -2, CanTake)
	_, _, squares = appendLegalSquare(squares, b, color, sq, 1, -2, CanTake)
	_, _, squares = appendLegalSquare(squares, b, color, sq, 2, -1, CanTake)

	return squares
}

func (p bishop) getLegalSquares(b board, sq square, color string) []square {
	return getLegalSquaresForBishop(b, sq, color)
}

func getLegalSquaresForBishop(b board, sq square, color string) []square {
	var squares []square
	var appended, willTakePiece bool
	var i, j int

	// Diagonally up and right from current position.
	i = sq.rank + 1
	j = fromFileStr(sq.file) + 1
	for {
		if i > BoardSize || j > BoardSize {
			break
		}
		appended, willTakePiece, squares = appendLegalSquare(squares, b, color, sq, i-sq.rank, j-fromFileStr(sq.file), CanTake)
		if !appended || (appended && willTakePiece) {
			break
		}

		i++
		j++
	}

	// Diagonally up and left from current position.
	i = sq.rank + 1
	j = fromFileStr(sq.file) - 1
	for {
		if i > BoardSize || j < 0 {
			break
		}
		appended, willTakePiece, squares = appendLegalSquare(squares, b, color, sq, i-sq.rank, j-fromFileStr(sq.file), CanTake)
		if !appended || (appended && willTakePiece) {
			break
		}

		i++
		j--
	}

	// Diagonally down and right from current position.
	i = sq.rank - 1
	j = fromFileStr(sq.file) + 1
	for {
		if i <= 0 || j >= BoardSize {
			break
		}
		appended, willTakePiece, squares = appendLegalSquare(squares, b, color, sq, i-sq.rank, j-fromFileStr(sq.file), CanTake)
		if !appended || (appended && willTakePiece) {
			break
		}

		i--
		j++
	}

	// Diagonally down and left from current position.
	i = sq.rank - 1
	j = fromFileStr(sq.file) - 1
	for {
		if i <= 0 || j < 0 {
			break
		}
		appended, willTakePiece, squares = appendLegalSquare(squares, b, color, sq, i-sq.rank, j-fromFileStr(sq.file), CanTake)
		if !appended || (appended && willTakePiece) {
			break
		}

		i--
		j--
	}

	return squares
}

func (p queen) getLegalSquares(b board, sq square, color string) []square {
	// Queen legal moves are effectively rook + bishop.
	squares := getLegalSquaresForRook(b, sq, color)
	squares = append(squares, getLegalSquaresForBishop(b, sq, color)...)
	return squares
}

func (p king) getLegalSquares(b board, sq square, color string) []square {
	var squares []square

	_, _, squares = appendLegalSquare(squares, b, color, sq, 1, 0, CanTake)
	_, _, squares = appendLegalSquare(squares, b, color, sq, 1, 1, CanTake)
	_, _, squares = appendLegalSquare(squares, b, color, sq, 0, 1, CanTake)
	_, _, squares = appendLegalSquare(squares, b, color, sq, -1, 1, CanTake)
	_, _, squares = appendLegalSquare(squares, b, color, sq, -1, 0, CanTake)
	_, _, squares = appendLegalSquare(squares, b, color, sq, -1, -1, CanTake)
	_, _, squares = appendLegalSquare(squares, b, color, sq, 0, -1, CanTake)
	_, _, squares = appendLegalSquare(squares, b, color, sq, 1, -1, CanTake)

	return squares
}

func appendLegalSquare(squares []square, b board, color string, sq square, rankOffset int, fileOffset int, tb takingBehavior) (bool, bool, []square) {

	if sq.rank+rankOffset <= 0 ||
		sq.rank+rankOffset > BoardSize ||
		fromFileStr(sq.file)+fileOffset < 0 ||
		fromFileStr(sq.file)+fileOffset >= BoardSize {
		return false, false, squares
	}

	newSquare := square{rank: sq.rank + rankOffset, file: toFileStr(fromFileStr(sq.file) + fileOffset)}

	if b.isSquareEmpty(newSquare) {
		// If piece has to take, can't move to empty square (e.g. pawn diagonals).
		if tb == MustTake {
			return false, false, squares
		}

		// Otherwise if OK to move to empty square.
		return true, false, append(squares, newSquare)
	}

	p, _ := b.getPieceAt(newSquare)
	if p.color != color && (tb == CanTake || tb == MustTake) {
		return true, true, append(squares, newSquare)
	}

	return false, false, squares
}

func (cp coloredPiece) String() string {
	return fmt.Sprintf("%s%s", cp.color, cp.piece.getName())
}
