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
	initPawns(b)
	initPieces(b, "B", 0)
	initPieces(b, "W", BoardSize-1)
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

func fromFileStr(s string) int {
	return strings.Index(Files, s)
}

func toFileStr(i int) string {
	return Files[i : i+1]
}
