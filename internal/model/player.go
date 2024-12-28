package model

import "fmt"

type Player int

const (
	WHITE Player = iota
	BLACK
)

var players = [...]string{"WHITE", "BLACK"}

var playerMap = map[string]Player{"WHITE": WHITE, "BLACK": BLACK}

func ParsePlayer(name string) (Player, error) {
	if pl, ok := playerMap[name]; ok {
		return pl, nil
	}
	return -1, fmt.Errorf("invalid player: %s", name)
}

func (p Player) String() string {
	if int(p) < len(players) {
		return players[p]
	}
	return "Invalid"
}
