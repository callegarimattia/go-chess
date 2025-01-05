package main

import (
	"os"

	"github.com/callegarimattia/go-chess/internal/model"
)

func main() {
	board := model.Board{}
	board.SetupPieces()
	board.Print(os.Stdout)
}
