package model

import (
	"fmt"
	"io"
)

type BitBoard uint64

func (bb BitBoard) getBit(square int) BitBoard {
	return bb & squareBB[square]
}

// func popBit(bb BitBoard, square int) BitBoard {
// 	return bb &^ squareBB[square]
// }
//
// func setBit(bb BitBoard, square int) BitBoard {
// 	return bb | squareBB[square]
// }
//
// func toggleBit(bb BitBoard, square int) BitBoard {
// 	return bb ^ squareBB[square]
// }

var squareBB [64]BitBoard //nolint: gochecknoglobals

func init() { //nolint:gochecknoinits
	for i := range 64 {
		squareBB[i] = 1 << i
	}
}

func (bb BitBoard) Print(writer io.Writer) {
	fmt.Fprintln(writer)

	for y := range 8 {
		for x := range 8 {
			square := 63 - ((y * 8) + x)
			if bit := bb & (1 << square); bit != 0 {
				fmt.Fprint(writer, "x")
			} else {
				fmt.Fprint(writer, ".")
			}

			fmt.Fprint(writer, " ")
		}

		fmt.Fprintln(writer)
	}

	fmt.Fprintln(writer)
}
