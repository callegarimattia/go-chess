package model

import (
	"fmt"
	"io"
)

const (
	secondRank       BitBoard = 0x00FF000000000000
	seventhRank      BitBoard = 0x000000000000FF00
	rightPawnCapture BitBoard = 0x8080808080808080
	leftPawnCapture  BitBoard = 0x0101010101010101
)

// A Board is represented by a 64bit unsigned integer for each piece, black and white.
// Little-Endian-Row-Last notation is used to map the squares to the 64bit uint.
type Board struct {
	whitePawns   BitBoard
	whiteKnights BitBoard
	whiteBishops BitBoard
	whiteRooks   BitBoard
	whiteQueens  BitBoard
	whiteKing    BitBoard

	blackPawns   BitBoard
	blackKnights BitBoard
	blackBishops BitBoard
	blackRooks   BitBoard
	blackQueens  BitBoard
	blackKing    BitBoard

	enemyPieces     BitBoard
	emptySquares    BitBoard
	attackedSquares BitBoard

	turn Player
}

// SetupPieces sets the pieces up on the board for a new game.
func (b *Board) SetupPieces() {
	b.setupWhitePieces()
	b.setupBlackPieces()
}

func (b *Board) setupBlackPieces() {
	b.blackPawns = 0x000000000000FF00
	b.blackKnights = 0x0000000000000042
	b.blackBishops = 0x0000000000000024
	b.blackRooks = 0x0000000000000081
	b.blackQueens = 0x0000000000000010
	b.blackKing = 0x0000000000000008
}

func (b *Board) setupWhitePieces() {
	b.whitePawns = 0x00FF000000000000
	b.whiteKnights = 0x4200000000000000
	b.whiteBishops = 0x2400000000000000
	b.whiteRooks = 0x8100000000000000
	b.whiteQueens = 0x1000000000000000
	b.whiteKing = 0x0800000000000000
}

// GetPieceBoards returns the bitboards represeting the full state of the board.
// They are returned in order, white pieces first, low to high value.
func (b *Board) GetPieceBoards() func(func(int, BitBoard) bool) {
	return func(yield func(int, BitBoard) bool) {
		bitboards := [...]BitBoard{
			b.whitePawns,
			b.whiteKnights,
			b.whiteBishops,
			b.whiteRooks,
			b.whiteQueens,
			b.whiteKing,

			b.blackPawns,
			b.blackKnights,
			b.blackBishops,
			b.blackRooks,
			b.blackQueens,
			b.blackKing,
		}
		for i := range bitboards {
			if !yield(i, bitboards[i]) {
				return
			}
		}
	}
}

func (b *Board) Print(writer io.Writer) {
	pieces := []rune("PNBRQKpnbrqk")

	fmt.Fprintln(writer, "  a b c d e f g h")

	for rank := range 8 {
		fmt.Fprint(writer, rank)

		for file := range 8 {
			square := 63 - ((rank * 8) + file)
			printed := false

			for i, bb := range b.GetPieceBoards() {
				if bb.getBit(square) > 0 {
					fmt.Fprintf(writer, " %c", pieces[i])

					printed = true

					break
				}
			}

			if !printed {
				fmt.Fprint(writer, " .")
			}
		}

		fmt.Fprintln(writer)
	}

	fmt.Fprintln(writer, "  a b c d e f g h")
	fmt.Fprintf(writer, "Turn: %s\n", b.turn)
}
