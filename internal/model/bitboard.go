package model

import "fmt"

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

var squareBB [64]BitBoard

func init() {
	for i := 0; i < 64; i++ {
		squareBB[i] = 1 << i
	}
}

func (b BitBoard) Print() {
	fmt.Println()
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			square := 63 - ((y * 8) + x)
			if bit := b & (1 << square); bit != 0 {
				fmt.Print("x")
			} else {
				fmt.Print(".")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println()
}
