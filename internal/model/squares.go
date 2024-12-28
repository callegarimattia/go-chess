package model

import "fmt"

type Square int

var squareMap = make(map[string]Square, len(squareNames))

func init() {
	for i, name := range squareNames {
		squareMap[name] = Square(i)
	}
}

func (s Square) String() string {
	if int(s) < len(squareNames) {
		return squareNames[s]
	}
	return "Invalid"
}

func ParseSquare(name string) (Square, error) {
	if sq, ok := squareMap[name]; ok {
		return sq, nil
	}
	return -1, fmt.Errorf("invalid square: %s", name)
}

var squareNames = [...]string{
	"H1", "G1", "F1", "E1", "D1", "C1", "B1", "A1",
	"H2", "G2", "F2", "E2", "D2", "C2", "B2", "A2",
	"H3", "G3", "F3", "E3", "D3", "C3", "B3", "A3",
	"H4", "G4", "F4", "E4", "D4", "C4", "B4", "A4",
	"H5", "G5", "F5", "E5", "D5", "C5", "B5", "A5",
	"H6", "G6", "F6", "E6", "D6", "C6", "B6", "A6",
	"H7", "G7", "F7", "E7", "D7", "C7", "B7", "A7",
	"H8", "G8", "F8", "E8", "D8", "C8", "B8", "A8",
}
