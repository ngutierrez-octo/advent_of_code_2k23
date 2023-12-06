package main

import (
	"adventofcode2k23/day_2"
	"adventofcode2k23/utils"
)

func main() {
	dataParser := utils.InputParser{}
	input, _ := dataParser.Execute("day_2/data/input_big.txt")
	day := day_2.Day2{}
	print(day.PlayPartTwo(input))
}
