package main

import "fmt"

type takingBehavior int

const (
	cannotTake takingBehavior = iota
	canTake
	mustTake
	mustTakeEnPassant
)

type piece interface {
	getName() string
	getLegalSquares(b board, sq square, color string, moved bool) []square
}

type gamePiece struct {
	piece
	color         string
	moved         bool
	numberOfMoves int
}

func (gp gamePiece) String() string {
	l := " "
	if gp.moved {
		l = "*"
	}
	return fmt.Sprintf("%s%s%s", gp.color, gp.piece.getName(), l)
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

func (p pawn) getLegalSquares(b board, sq square, color string, moved bool) []square {
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
		appended, _, squares = appendLegalSquare(squares, b, p, color, sq, 1*direction, 0, cannotTake)
	}

	// Double move forward - allowed if single move was allowed, and on second rank.
	if appended && sq.rank == secondRank {
		appended, _, squares = appendLegalSquare(squares, b, p, color, sq, 2*direction, 0, cannotTake)
	}

	// Diagonal taking moves
	_, _, squares = appendLegalSquare(squares, b, p, color, sq, 1*direction, 1, mustTake)
	_, _, squares = appendLegalSquare(squares, b, p, color, sq, 1*direction, -1, mustTake)

	// En passant taking moves
	_, _, squares = appendLegalSquare(squares, b, p, color, sq, 1*direction, 1, mustTakeEnPassant)
	_, _, squares = appendLegalSquare(squares, b, p, color, sq, 1*direction, -1, mustTakeEnPassant)

	return squares
}

func (p rook) getLegalSquares(b board, sq square, color string, moved bool) []square {
	return getLegalSquaresForRook(b, p, sq, color)
}

func getLegalSquaresForRook(b board, p piece, sq square, color string) []square {
	var squares []square
	var appended, willTakePiece bool

	// Vertically up from current position.
	for i := sq.rank + 1; i < BoardSize; i++ {
		appended, willTakePiece, squares = appendLegalSquare(squares, b, p, color, sq, i-sq.rank, 0, canTake)
		if !appended || (appended && willTakePiece) {
			break
		}
	}

	// Vertically down from current position.
	for i := sq.rank - 1; i > 0; i-- {
		appended, willTakePiece, squares = appendLegalSquare(squares, b, p, color, sq, i-sq.rank, 0, canTake)
		if !appended || (appended && willTakePiece) {
			break
		}
	}

	// Horizonally right from current position.
	fileNumber := fromFileStr(sq.file)
	for i := fileNumber + 1; i < BoardSize; i++ {
		appended, willTakePiece, squares = appendLegalSquare(squares, b, p, color, sq, 0, i-fileNumber, canTake)
		if !appended || (appended && willTakePiece) {
			break
		}
	}

	// Horizonally left from current position.
	for i := fileNumber - 1; i >= 0; i-- {
		appended, willTakePiece, squares = appendLegalSquare(squares, b, p, color, sq, 0, i-fileNumber, canTake)
		if !appended || (appended && willTakePiece) {
			break
		}
	}

	return squares
}

func (p knight) getLegalSquares(b board, sq square, color string, moved bool) []square {
	var squares []square

	_, _, squares = appendLegalSquare(squares, b, p, color, sq, 2, 1, canTake)
	_, _, squares = appendLegalSquare(squares, b, p, color, sq, 1, 2, canTake)
	_, _, squares = appendLegalSquare(squares, b, p, color, sq, -1, 2, canTake)
	_, _, squares = appendLegalSquare(squares, b, p, color, sq, -2, 1, canTake)
	_, _, squares = appendLegalSquare(squares, b, p, color, sq, -2, -1, canTake)
	_, _, squares = appendLegalSquare(squares, b, p, color, sq, -1, -2, canTake)
	_, _, squares = appendLegalSquare(squares, b, p, color, sq, 1, -2, canTake)
	_, _, squares = appendLegalSquare(squares, b, p, color, sq, 2, -1, canTake)

	return squares
}

func (p bishop) getLegalSquares(b board, sq square, color string, moved bool) []square {
	return getLegalSquaresForBishop(b, p, sq, color)
}

func getLegalSquaresForBishop(b board, p piece, sq square, color string) []square {
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
		appended, willTakePiece, squares = appendLegalSquare(squares, b, p, color, sq, i-sq.rank, j-fromFileStr(sq.file), canTake)
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
		appended, willTakePiece, squares = appendLegalSquare(squares, b, p, color, sq, i-sq.rank, j-fromFileStr(sq.file), canTake)
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
		appended, willTakePiece, squares = appendLegalSquare(squares, b, p, color, sq, i-sq.rank, j-fromFileStr(sq.file), canTake)
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
		appended, willTakePiece, squares = appendLegalSquare(squares, b, p, color, sq, i-sq.rank, j-fromFileStr(sq.file), canTake)
		if !appended || (appended && willTakePiece) {
			break
		}

		i--
		j--
	}

	return squares
}

func (p queen) getLegalSquares(b board, sq square, color string, moved bool) []square {
	// Queen legal moves are effectively rook + bishop.
	squares := getLegalSquaresForRook(b, p, sq, color)
	squares = append(squares, getLegalSquaresForBishop(b, p, sq, color)...)
	return squares
}

func (p king) getLegalSquares(b board, sq square, color string, moved bool) []square {
	var squares []square

	// Single square moves
	_, _, squares = appendLegalSquare(squares, b, p, color, sq, 1, 0, canTake)
	_, _, squares = appendLegalSquare(squares, b, p, color, sq, 1, 1, canTake)
	_, _, squares = appendLegalSquare(squares, b, p, color, sq, 0, 1, canTake)
	_, _, squares = appendLegalSquare(squares, b, p, color, sq, -1, 1, canTake)
	_, _, squares = appendLegalSquare(squares, b, p, color, sq, -1, 0, canTake)
	_, _, squares = appendLegalSquare(squares, b, p, color, sq, -1, -1, canTake)
	_, _, squares = appendLegalSquare(squares, b, p, color, sq, 0, -1, canTake)
	_, _, squares = appendLegalSquare(squares, b, p, color, sq, 1, -1, canTake)

	// Castling
	// - king that has moved cannot castle
	if moved {
		return squares
	}

	rookSquares := getRookSquaresForKing(b, sq)
	if canCastle(b, sq, rookSquares[0], color) {
		_, _, squares = appendLegalSquare(squares, b, p, color, sq, 0, -2, cannotTake)
	}
	if canCastle(b, sq, rookSquares[1], color) {
		_, _, squares = appendLegalSquare(squares, b, p, color, sq, 0, 2, cannotTake)
	}

	return squares
}

func getRookSquaresForKing(b board, kingSquare square) [2]square {
	result := [2]square{}
	result[0] = square{file: toFileStr(0), rank: kingSquare.rank}
	result[1] = square{file: toFileStr(BoardSize - 1), rank: kingSquare.rank}
	return result
}

func canCastle(b board, kingSquare square, rookSquare square, color string) bool {
	// Must have an unmoved rook to castle with.
	if b.isSquareEmpty(rookSquare) {
		return false
	}

	piece, _ := b.getPieceAt(rookSquare)
	if piece.getName() != "R" || piece.moved {
		return false
	}

	// Can't be any blocking pieces.
	if !b.areEmptySquaresBetween(kingSquare, rookSquare) {
		return false
	}

	// Can't castle through check
	for _, sq := range getSquaresBetween(kingSquare, rookSquare) {
		isSquareEnPrise, _ := isSquareEnPrise(b, sq, color)
		if isSquareEnPrise {
			return false
		}
	}

	return true
}

func appendLegalSquare(squares []square, b board, p piece, color string, sq square, rankOffset int, fileOffset int, tb takingBehavior) (bool, bool, []square) {

	if sq.rank+rankOffset <= 0 ||
		sq.rank+rankOffset > BoardSize ||
		fromFileStr(sq.file)+fileOffset < 0 ||
		fromFileStr(sq.file)+fileOffset >= BoardSize {
		return false, false, squares
	}

	newSquare := square{rank: sq.rank + rankOffset, file: toFileStr(fromFileStr(sq.file) + fileOffset)}

	if b.isSquareEmpty(newSquare) {
		// If piece has to take, can't move to empty square (e.g. pawn diagonals).
		if tb == mustTake {
			return false, false, squares
		}

		// Can move to square en passant only if piece is a pawn, it's moving to the 6th rank,
		// and an opposing pawn that's made the last move is on the 5th rank.
		if tb == mustTakeEnPassant {
			if p.getName() != "P" ||
				!((color == "W" && sq.rank == 5) || (color == "B" && sq.rank == 4)) {
				return false, false, squares
			}

			enPassantTakingSquare := square{rank: sq.rank, file: toFileStr(fromFileStr(sq.file) + fileOffset)}
			if b.isSquareEmpty(enPassantTakingSquare) {
				return false, false, squares
			}

			pieceToTakeEnPassant, _ := b.getPieceAt(enPassantTakingSquare)
			if pieceToTakeEnPassant.getName() != "P" || pieceToTakeEnPassant.numberOfMoves != 1 {
				return false, false, squares
			}
		}

		// Otherwise if OK to move to empty square.
		return true, false, append(squares, newSquare)
	}

	pieceToTake, _ := b.getPieceAt(newSquare)

	// Can only move to occupied square if piece is of opposite colour and taking allowed.
	if pieceToTake.color != color && (tb == canTake || tb == mustTake) {
		return true, true, append(squares, newSquare)
	}

	return false, false, squares
}
