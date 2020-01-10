package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	board := board{}
	board.init()
	board.print()

	color := "W"
	kingInCheck := false

	reader := bufio.NewReader(os.Stdin)
	for {
		if board.isKingInCheckMate(color) {
			fmt.Printf("The %s king is in checkmate. %s wins!.\n", color, switchColor(color))
			break
		}

		kingInCheck, _ = board.isKingInCheck(color)
		if kingInCheck {
			fmt.Printf("The %s king is in check!\n", color)
		}

		fmt.Printf("Select piece (%s): ", color)
		fromInput, _ := reader.ReadString('\n')
		fromSquare, err := getSquareFromInput(fromInput)
		if err != nil {
			fmt.Println(err)
			continue
		}

		piece, err := board.getPieceAt(fromSquare)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if piece.color != color {
			fmt.Printf("Piece isn't of the correct colour (%s)\n", color)
			continue
		}

		fmt.Printf("Enter destination square: ")
		toInput, _ := reader.ReadString('\n')
		toSquare, err := getSquareFromInput(toInput)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if !isMoveLegal(board, piece, fromSquare, toSquare) {
			fmt.Println("Not a legal move.")
			continue
		}

		kingInCheck, _ = board.isKingInCheck(color)
		if kingInCheck {
			fmt.Println("Not a legal move (your king would be in check).")
			continue
		}

		board.movePiece(fromSquare, toSquare)
		board.print()

		color = switchColor(color)
	}
}

func getSquareFromInput(entry string) (square, error) {
	if utf8.RuneCountInString(entry) != 4 {
		return square{}, errors.New("Entry not valid (must be 2 characters).")
	}

	// TODO: more validation

	file := strings.ToUpper(entry[0:1])
	rank, _ := strconv.Atoi(entry[1:2])

	return square{file, rank}, nil
}

func isMoveLegal(b board, p coloredPiece, fromSquare square, toSquare square) bool {
	legalSquares := p.getLegalSquares(b, fromSquare, p.color)
	for _, sq := range legalSquares {
		if sq.rank == toSquare.rank && sq.file == toSquare.file {
			return true
		}
	}

	return false
}

func switchColor(color string) string {
	if color == "W" {
		return "B"
	}

	return "W"
}
