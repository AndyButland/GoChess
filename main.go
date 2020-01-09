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

	turn := "W"

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Select piece (%s): ", turn)
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

		if piece.color != turn {
			fmt.Printf("Piece isn't of the correct colour (%s)\n", turn)
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

		board.movePiece(fromSquare, toSquare)
		board.print()

		turn = switchTurn(turn)
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

func switchTurn(turn string) string {
	if turn == "W" {
		return "B"
	}

	return "W"
}
