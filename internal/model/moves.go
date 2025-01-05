package model

func generatePawnMovesWhite(pawns, emptySquares, enemyPieces BitBoard) BitBoard {
	single := singlePawnPushWhite(pawns, emptySquares)
	double := doublePawnPushWhite(pawns, emptySquares)
	captures := capturesPawnsWhite(pawns, enemyPieces)
	// enpassant := enpassantWhite(pawns, enpassantTarget)

	return single | double | captures
}

func singlePawnPushWhite(pawns, emptySquares BitBoard) BitBoard {
	return (pawns >> 8) & emptySquares
}

func doublePawnPushWhite(pawns, emptySquares BitBoard) BitBoard {
	singlePush := singlePawnPushWhite(pawns, emptySquares)

	return (singlePush >> 8) & emptySquares & secondRank
}

func capturesPawnsWhite(pawns, enemyPieces BitBoard) BitBoard {
	leftCapture := (pawns >> 9) & enemyPieces & leftPawnCapture
	rightCapture := (pawns >> 7) & enemyPieces & rightPawnCapture

	return leftCapture | rightCapture
}

// TODO: implement enpassant
// func enpassantWhite(pawns, enpassantTarget BitBoard) BitBoard {
//    return
// }

func singlePawnPushBlack(pawns, emptySquares BitBoard) BitBoard {
	return (pawns << 8) & emptySquares
}

func doublePawnPushBlack(pawns, emptySquares BitBoard) BitBoard {
	singlePush := singlePawnPushBlack(pawns, emptySquares)

	return (singlePush << 8) & emptySquares & seventhRank
}

func capturesPawnsBlack(pawns, enemyPieces BitBoard) BitBoard {
	leftCapture := (pawns << 7) & enemyPieces & leftPawnCapture
	rightCapture := (pawns << 9) & enemyPieces & rightPawnCapture

	return leftCapture | rightCapture
}
