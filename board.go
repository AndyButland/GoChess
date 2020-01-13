package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

const BoardSize = 8
const Files = "ABCDEFGH"

type square struct {
	file string
	rank int
}

type board [BoardSize][BoardSize]gamePiece

func (b *board) init() {
	b.clear()
	initPawns(b)
	initPieces(b, "B", 0)
	initPieces(b, "W", BoardSize-1)
}

func (b *board) clear() {
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			b.setSquareEmpty(i, j)
		}
	}
}

func initPawns(b *board) {
	for i := 0; i < BoardSize; i++ {
		(*b)[1][i] = gamePiece{color: "B", piece: pawn{}}
		(*b)[6][i] = gamePiece{color: "W", piece: pawn{}}
	}
}

func initPieces(b *board, color string, row int) {
	(*b)[row][0] = gamePiece{color: color, piece: rook{}}
	(*b)[row][1] = gamePiece{color: color, piece: knight{}}
	(*b)[row][2] = gamePiece{color: color, piece: bishop{}}
	(*b)[row][3] = gamePiece{color: color, piece: queen{}}
	(*b)[row][4] = gamePiece{color: color, piece: king{}}
	(*b)[row][5] = gamePiece{color: color, piece: bishop{}}
	(*b)[row][6] = gamePiece{color: color, piece: knight{}}
	(*b)[row][7] = gamePiece{color: color, piece: rook{}}
}

func (b board) isSquareEmpty(sq square) bool {
	return b.isRowColEmpty(getRowColForSquare(sq))
}

func (b board) isRowColEmpty(row int, col int) bool {
	return (gamePiece{}) == b[row][col]
}

func (b board) getPieceAt(sq square) (gamePiece, error) {
	if b.isSquareEmpty(sq) {
		return gamePiece{}, errors.New("No piece found at square")
	}

	row, col := getRowColForSquare(sq)
	return b[row][col], nil
}

func areSquaresEqual(sq1 square, sq2 square) bool {
	return sq1.file == sq2.file && sq1.rank == sq2.rank
}

func areSquaresAdjacent(sq1 square, sq2 square) bool {
	return math.Abs(float64(sq1.rank-sq2.rank)) <= 1 &&
		math.Abs(float64(fromFileStr(sq1.file)-fromFileStr(sq2.file))) <= 1
}

func getSquaresBetween(sq1 square, sq2 square) []square {
	var squares []square
	if areSquaresEqual(sq1, sq2) || areSquaresAdjacent(sq1, sq2) {
		return squares
	}

	minRank := minOf(sq1.rank, sq2.rank)
	maxRank := maxOf(sq1.rank, sq2.rank)
	minFile := minOf(fromFileStr(sq1.file), fromFileStr(sq2.file))
	maxFile := maxOf(fromFileStr(sq1.file), fromFileStr(sq2.file))

	if sq1.file == sq2.file {
		// Vertical squares between
		for i := minRank + 1; i < maxRank; i++ {
			squares = append(squares, square{file: sq1.file, rank: i})
		}
	} else if sq1.rank == sq2.rank {
		// Horizontal squares between
		for i := minFile + 1; i < maxFile; i++ {
			squares = append(squares, square{file: toFileStr(i), rank: sq1.rank})
		}
	} else if maxRank-minRank == maxFile-minFile {
		// Diagonal squares between
		var leftRightDirection int
		if sq2.file > sq1.file {
			leftRightDirection = 1
		} else {
			leftRightDirection = -1
		}
		count := 1
		for i := minRank + 1; i < maxRank; i++ {
			squares = append(squares, square{file: toFileStr(fromFileStr(sq1.file) + count*leftRightDirection), rank: i})
			count++
		}
	}

	return squares
}

func (b board) areEmptySquaresBetween(sq1 square, sq2 square) bool {
	squares := getSquaresBetween(sq1, sq2)
	if len(squares) == 0 {
		return false
	}

	for _, sq := range squares {
		if !b.isSquareEmpty(sq) {
			return false
		}
	}

	return true
}

func (b *board) movePiece(fromSquare square, toSquare square) {
	fromRow, fromCol := getRowColForSquare(fromSquare)
	toRow, toCol := getRowColForSquare(toSquare)
	piece := (*b)[fromRow][fromCol]
	(*b)[toRow][toCol] = piece
	piece.moved = true
	b.setSquareEmpty(fromRow, fromCol)

	if isCastling(piece, fromCol, toCol) {
		moveCastledRook(b, fromRow, toCol)
	}
}

func isCastling(gp gamePiece, fromCol int, toCol int) bool {
	return gp.getName() == "K" && math.Abs(float64(fromCol)-float64(toCol)) == 2
}

func moveCastledRook(b *board, row int, kingCol int) {
	var currentSquare square
	var newSquare square
	if kingCol > BoardSize/2 {
		currentSquare = square{rank: BoardSize - row, file: toFileStr(BoardSize - 1)}
		newSquare = square{rank: BoardSize - row, file: toFileStr(kingCol - 1)}
	} else {
		currentSquare = square{rank: BoardSize - row, file: toFileStr(0)}
		newSquare = square{rank: BoardSize - row, file: toFileStr(kingCol + 1)}
	}

	b.movePiece(currentSquare, newSquare)
}

func (b *board) setSquareEmpty(row int, col int) {
	(*b)[row][col] = gamePiece{}
}

func (b board) isKingInCheck(color string) (bool, []square) {
	kingSquare, _ := b.getSquareForPiece(color, "K")

	// To determine if king is in check, we can more generally check if the piece can be "taken",
	// i.e. is en prise.
	return isSquareEnPrise(b, kingSquare, color)
}

func (b board) isKingInCheckMate(color string) (bool, string) {
	// If king not in check, can't be in check-mate.
	kingInCheck, checkingSquares := b.isKingInCheck(color)
	if !kingInCheck {
		return false, "Not in check"
	}

	// King is in check.  It won't be check-mate though, if:
	kingSquare, _ := b.getSquareForPiece(color, "K")
	king, _ := b.getPieceAt(kingSquare)
	var opponentColor string
	if color == "W" {
		opponentColor = "B"
	} else {
		opponentColor = "W"
	}

	// - king has legal moves, and at least one moves it out of check
	if kingCanMoveOutOfCheck(b, kingSquare, king, color) {
		return false, "King can move out of check"
	}

	// - OR any piece has a legal move that takes the checking piece (if there's only one of them),
	//   unless the only piece that can take is the king, which would then still be in check)
	if len(checkingSquares) == 1 {
		isSquareEnPrise, takingSquares := isSquareEnPrise(b, checkingSquares[0], opponentColor)
		if isSquareEnPrise && (len(takingSquares) > 1 || !takingPieceIsKingMovingToCheck(b, takingSquares[0], checkingSquares[0])) {
			return false, "Checking piece can be taken"
		}
	}

	// - OR check can be blocked by a piece
	// -- if more than one checking piece, can't be blocked
	if len(checkingSquares) > 1 {
		return true, "In check, and more than one checking piece so can't block or take"
	}

	// -- knights can't be blocked
	checkingPiece, _ := b.getPieceAt(checkingSquares[0])
	if checkingPiece.getName() == "N" {
		return true, "In check, can't take and checking piece is a knight so can't be blocked"
	}

	// -- if piece is adjacent, can't be blocked
	if math.Abs(float64(checkingSquares[0].rank-kingSquare.rank)) <= 1 &&
		math.Abs(float64(fromFileStr(checkingSquares[0].file)-fromFileStr(kingSquare.file))) <= 1 {
		return true, "In check, can't take and checking piece is adjacent so can't be blocked"
	}

	// -- can block if any piece has a legal move that intercepts the vertical, horizontal or
	//    diagonal line between the single checking piece and the king
	for _, squareBetween := range getSquaresBetween(kingSquare, checkingSquares[0]) {
		for i := 0; i < BoardSize; i++ {
			for j := 0; j < BoardSize; j++ {
				if !b.isRowColEmpty(i, j) {
					piece := b[i][j]
					if piece.color == color && piece.getName() != "K" {
						square := getSquareForRowCol(i, j)
						for _, legalSquare := range piece.getLegalSquares(b, square, piece.color, piece.moved) {
							if areSquaresEqual(squareBetween, legalSquare) {
								return false, fmt.Sprintf("In check, can't take but piece %s on %v can block", piece.getName(), square)
							}
						}
					}
				}
			}
		}
	}

	return true, "In check, can't take or block"
}

func isSquareEnPrise(b board, pieceSquare square, color string) (bool, []square) {
	// To determine if a square is en prise is in check, we look at legal moves for all
	// the opponent's pieces, and if they include the piece, it's en prise.
	var takingSquares []square
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if !b.isRowColEmpty(i, j) {
				piece := b[i][j]
				if piece.color != color {
					square := getSquareForRowCol(i, j)
					legalSquares := piece.getLegalSquares(b, square, piece.color, piece.moved)
					for _, sq := range legalSquares {
						if sq.rank == pieceSquare.rank && sq.file == pieceSquare.file {
							takingSquares = append(takingSquares, square)
							break
						}
					}
				}
			}
		}
	}

	if (len(takingSquares)) > 0 {
		return true, takingSquares
	}

	return false, takingSquares
}

func kingCanMoveOutOfCheck(b board, kingSquare square, k gamePiece, color string) bool {
	legalSquares := k.getLegalSquares(b, kingSquare, color, k.moved)
	if len(legalSquares) > 0 {
		for _, sq := range legalSquares {
			tempBoard := b
			tempBoard.movePiece(kingSquare, sq)
			movedKingInCheck, _ := tempBoard.isKingInCheck(color)
			if !movedKingInCheck {
				return true
			}
		}
	}

	return false
}

func takingPieceIsKingMovingToCheck(b board, fromSquare square, toSquare square) bool {
	takingPiece, _ := b.getPieceAt(fromSquare)
	if takingPiece.getName() != "K" {
		return false
	}

	tempBoard := b
	tempBoard.movePiece(fromSquare, toSquare)
	isKingInCheck, _ := b.isKingInCheck(takingPiece.color)
	return isKingInCheck
}

func (b board) getSquareForPiece(color string, name string) (square, error) {
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if !b.isRowColEmpty(i, j) {
				piece := b[i][j]
				if piece.color == color && piece.getName() == "K" {
					return getSquareForRowCol(i, j), nil
				}
			}
		}
	}

	return square{}, fmt.Errorf("Piece %s%s not found", color, name)
}

func (b board) print() {
	fmt.Println()
	printRankSeparator(b)
	for i := 0; i < BoardSize; i++ {
		fmt.Printf("%d ", BoardSize-i)
		for j := 0; j < BoardSize; j++ {
			if b.isRowColEmpty(i, j) {
				fmt.Printf("|  ")
			} else {
				fmt.Printf("|%2s", b[i][j])
			}
		}

		fmt.Printf("|\n")
		printRankSeparator(b)
	}

	fmt.Printf("   ")
	for i := 0; i < BoardSize; i++ {
		fmt.Printf("%s  ", toFileStr(i))
	}

	fmt.Println()
	fmt.Println()
}

func printRankSeparator(b board) {
	fmt.Println("  " + strings.Repeat("-", BoardSize*3+1))
}

func getRowColForSquare(sq square) (row int, col int) {
	return BoardSize - sq.rank, fromFileStr(sq.file)
}

func getSquareForRowCol(row int, col int) square {
	return square{file: toFileStr(col), rank: BoardSize - row}
}

func fromFileStr(s string) int {
	return strings.Index(Files, s)
}

func toFileStr(i int) string {
	return Files[i : i+1]
}
