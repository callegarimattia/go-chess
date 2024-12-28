package model

import "fmt"

type Piece int

var pieceNames = [...]string{
	"PAWN_W", "KNIGHT_W", "BISHOP_W", "ROOK_W", "QUEEN_W", "KING_W",
	"PAWN_B", "KNIGHT_B", "BISHOP_B", "ROOK_B", "QUEEN_B", "KING_B",
}

var pieceMap = make(map[string]Piece, len(pieceNames))

func init() {
	for i, name := range pieceNames {
		pieceMap[name] = Piece(i)
	}
}

func (p Piece) String() string {
	if int(p) < len(pieceNames) {
		return pieceNames[p]
	}
	return "Invalid"
}

func ParsePiece(name string) (Piece, error) {
	if piece, ok := pieceMap[name]; ok {
		return piece, nil
	}
	return -1, fmt.Errorf("invalid piece: %s", name)
}
