package day_2

type Day2 struct {
}

func (d Day2) PlayPartOne(games []string) int {
	score := 0
	for gameId, game := range games {
		gameParser := GameParser{}
		gameParsed := gameParser.Execute(game)

		if checkLegalGame(gameParsed) {
			score += gameId + 1
		}
	}
	return score
}

func (d Day2) PlayPartTwo(games []string) int {
	gameParser := GameParser{}
	score := 0
	for _, game := range games {
		gameParsed := gameParser.Execute(game)
		score += getPower(gameParsed)

	}
	return score

}

func getPower(game []map[string]int) int {
	power := 1
	qtyOfCubesMinRequired := map[string]int{
		"red":   0,
		"blue":  0,
		"green": 0,
	}

	for _, move := range game {
		for color, qty := range move {
			if qty > qtyOfCubesMinRequired[color] {
				qtyOfCubesMinRequired[color] = qty
			}
		}
	}
	for _, qty := range qtyOfCubesMinRequired {
		power *= qty
	}
	return power

}

func checkLegalGame(game []map[string]int) bool {
	for _, move := range game {
		if !checkLegalMove(move) {
			return false
		}
	}
	return true
}

func checkLegalMove(move map[string]int) bool {
	cubeStock := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for color, qty := range move {
		if qty > cubeStock[color] {
			return false
		}
	}

	return true
}
