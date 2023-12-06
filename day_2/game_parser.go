package day_2

import (
	"strconv"
	"strings"
)

type GameParser struct {
	games []string
}

func (gameParser *GameParser) RemovePrefix(game string) string {
	indexEndOfPrefix := strings.Index(game, ": ")
	return game[indexEndOfPrefix+2:]
}

func (gameParser *GameParser) parseSingleMove(move string) map[string]int {
	// "1 blue, 2 green"
	moveParsed := make(map[string]int)
	cubes := strings.Split(move, ", ")

	for _, cube := range cubes {
		cubeParsed := strings.Split(cube, " ")
		moveParsed[cubeParsed[1]], _ = strconv.Atoi(cubeParsed[0])

	}
	return moveParsed
}

func (gameParser *GameParser) Execute(game string) []map[string]int {
	var gameParsed []map[string]int
	moves := strings.Split(gameParser.RemovePrefix(game), "; ")
	for _, move := range moves {
		gameParsed = append(gameParsed, gameParser.parseSingleMove(move))
	}
	return gameParsed

}
