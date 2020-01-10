package main

import (
	"errors"
	"fmt"
	"strings"
)

const BoardSize = 8
const Files = "ABCDEFGH"

type square struct {
	file string
	rank int
}

type board [BoardSize][BoardSize]coloredPiece

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
		(*b)[1][i] = coloredPiece{color: "B", piece: pawn{}}
		(*b)[6][i] = coloredPiece{color: "W", piece: pawn{}}
	}
}

func initPieces(b *board, color string, row int) {
	(*b)[row][0] = coloredPiece{color: color, piece: rook{}}
	(*b)[row][1] = coloredPiece{color: color, piece: knight{}}
	(*b)[row][2] = coloredPiece{color: color, piece: bishop{}}
	(*b)[row][3] = coloredPiece{color: color, piece: queen{}}
	(*b)[row][4] = coloredPiece{color: color, piece: king{}}
	(*b)[row][5] = coloredPiece{color: color, piece: bishop{}}
	(*b)[row][6] = coloredPiece{color: color, piece: knight{}}
	(*b)[row][7] = coloredPiece{color: color, piece: rook{}}
}

func (b board) isSquareEmpty(sq square) bool {
	return b.isRowColEmpty(getRowColForSquare(sq))
}

func (b board) isRowColEmpty(row int, col int) bool {
	return (coloredPiece{}) == b[row][col]
}

func (b board) getPieceAt(sq square) (coloredPiece, error) {
	if b.isSquareEmpty(sq) {
		return coloredPiece{}, errors.New("No piece found at square")
	}

	row, col := getRowColForSquare(sq)
	return b[row][col], nil
}

func (b *board) movePiece(fromSquare square, toSquare square) {
	fromRow, fromCol := getRowColForSquare(fromSquare)
	toRow, toCol := getRowColForSquare(toSquare)
	(*b)[toRow][toCol] = (*b)[fromRow][fromCol]
	b.setSquareEmpty(fromRow, fromCol)
}

func (b *board) setSquareEmpty(row int, col int) {
	(*b)[row][col] = coloredPiece{}
}

func (b board) isKingInCheck(color string) (bool, []square) {
	kingSquare, _ := b.getSquareForPiece(color, "K")

	// To determine if king is in check, we can more generally check if the piece can be "taken",
	// i.e. is en prise.
	return isSquareEnPrise(b, kingSquare, color)
}

func (b board) isKingInCheckMate(color string) bool {
	// If king not in check, can't be in check-mate.
	kingInCheck, checkingSquares := b.isKingInCheck(color)
	if !kingInCheck {
		return false
	}

	// King is in check.  It won't be check-mate though, if:
	kingSquare, _ := b.getSquareForPiece(color, "K")
	king, _ := b.getPieceAt(kingSquare)

	// - king has legal moves, and at least one moves it out of check
	if kingCanMoveOutOfCheck(b, kingSquare, king, color) {
		return false
	}

	// - OR any piece has a legal move that takes the checking piece (if there's only one of them)
	if len(checkingSquares) == 1 {
		var opponentColor string
		if color == "W" {
			opponentColor = "B"
		} else {
			opponentColor = "W"
		}
		isSquareEnPrise, _ := isSquareEnPrise(b, checkingSquares[0], opponentColor)
		if isSquareEnPrise {
			fmt.Println("Checking piece can be taken")
			return false
		}
	}

	// TODO:
	// - OR check can be blocked by a piece

	return true
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
					legalSquares := piece.getLegalSquares(b, square, piece.color)
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

func kingCanMoveOutOfCheck(b board, kingSquare square, k coloredPiece, color string) bool {
	legalSquares := k.getLegalSquares(b, kingSquare, color)
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
