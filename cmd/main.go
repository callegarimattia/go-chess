package main

import (
	"github.com/callegarimattia/go-chess/internal/model"
)

func main() {
	board := model.Board{}
	board.SetupPieces()
	board.Print()
}
