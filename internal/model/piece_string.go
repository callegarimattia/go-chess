// Code generated by "stringer -type Piece,Player,Square"; DO NOT EDIT.

package model

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PAWN_W-0]
	_ = x[KNIGHT_W-1]
	_ = x[BISHOP_W-2]
	_ = x[ROOK_W-3]
	_ = x[QUEEN_W-4]
	_ = x[KING_W-5]
	_ = x[PAWN_B-6]
	_ = x[KNIGHT_B-7]
	_ = x[BISHOP_B-8]
	_ = x[ROOK_B-9]
	_ = x[QUEEN_B-10]
	_ = x[KING_B-11]
}

const _Piece_name = "PAWN_WKNIGHT_WBISHOP_WROOK_WQUEEN_WKING_WPAWN_BKNIGHT_BBISHOP_BROOK_BQUEEN_BKING_B"

var _Piece_index = [...]uint8{0, 6, 14, 22, 28, 35, 41, 47, 55, 63, 69, 76, 82}

func (i Piece) String() string {
	if i < 0 || i >= Piece(len(_Piece_index)-1) {
		return "Piece(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Piece_name[_Piece_index[i]:_Piece_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[WHITE-0]
	_ = x[BLACK-1]
}

const _Player_name = "WHITEBLACK"

var _Player_index = [...]uint8{0, 5, 10}

func (i Player) String() string {
	if i < 0 || i >= Player(len(_Player_index)-1) {
		return "Player(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Player_name[_Player_index[i]:_Player_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[H1-0]
	_ = x[G1-1]
	_ = x[F1-2]
	_ = x[E1-3]
	_ = x[D1-4]
	_ = x[C1-5]
	_ = x[B1-6]
	_ = x[A1-7]
	_ = x[H2-8]
	_ = x[G2-9]
	_ = x[F2-10]
	_ = x[E2-11]
	_ = x[D2-12]
	_ = x[C2-13]
	_ = x[B2-14]
	_ = x[A2-15]
	_ = x[H3-16]
	_ = x[G3-17]
	_ = x[F3-18]
	_ = x[E3-19]
	_ = x[D3-20]
	_ = x[C3-21]
	_ = x[B3-22]
	_ = x[A3-23]
	_ = x[H4-24]
	_ = x[G4-25]
	_ = x[F4-26]
	_ = x[E4-27]
	_ = x[D4-28]
	_ = x[C4-29]
	_ = x[B4-30]
	_ = x[A4-31]
	_ = x[H5-32]
	_ = x[G5-33]
	_ = x[F5-34]
	_ = x[E5-35]
	_ = x[D5-36]
	_ = x[C5-37]
	_ = x[B5-38]
	_ = x[A5-39]
	_ = x[H6-40]
	_ = x[G6-41]
	_ = x[F6-42]
	_ = x[E6-43]
	_ = x[D6-44]
	_ = x[C6-45]
	_ = x[B6-46]
	_ = x[A6-47]
	_ = x[H7-48]
	_ = x[G7-49]
	_ = x[F7-50]
	_ = x[E7-51]
	_ = x[D7-52]
	_ = x[C7-53]
	_ = x[B7-54]
	_ = x[A7-55]
	_ = x[H8-56]
	_ = x[G8-57]
	_ = x[F8-58]
	_ = x[E8-59]
	_ = x[D8-60]
	_ = x[C8-61]
	_ = x[B8-62]
	_ = x[A8-63]
}

const _Square_name = "H1G1F1E1D1C1B1A1H2G2F2E2D2C2B2A2H3G3F3E3D3C3B3A3H4G4F4E4D4C4B4A4H5G5F5E5D5C5B5A5H6G6F6E6D6C6B6A6H7G7F7E7D7C7B7A7H8G8F8E8D8C8B8A8"

var _Square_index = [...]uint8{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128}

func (i Square) String() string {
	if i < 0 || i >= Square(len(_Square_index)-1) {
		return "Square(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Square_name[_Square_index[i]:_Square_index[i+1]]
}
