package day_2

import (
	"reflect"
	"testing"
)

func TestRemovePrefix(t *testing.T) {
	input := "Game 2: 1 blue, 2 green"

	output := "1 blue, 2 green"

	gameParser := GameParser{}
	if !reflect.DeepEqual(gameParser.RemovePrefix(input), output) {
		println(gameParser.RemovePrefix(input))
		t.Errorf("Error trying to remove prefix")
	}
}

func TestParseSingleMove(t *testing.T) {
	input := "1 blue, 2 green"

	output := map[string]int{
		"blue":  1,
		"green": 2,
	}
	gameParser := GameParser{}
	outputObtained := gameParser.parseSingleMove(input)

	if !reflect.DeepEqual(outputObtained, output) {
		t.Errorf("Error trying to parse a single move")
	}
}

func TestParseGame(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

	output := []map[string]int{
		{
			"blue": 3,
			"red":  4,
		},
		{
			"red":   1,
			"green": 2,
			"blue":  6,
		},
		{
			"green": 2,
		},
	}
	gameParser := GameParser{}
	outputObtained := gameParser.Execute(input)

	if !reflect.DeepEqual(outputObtained, output) {
		t.Errorf("Error trying to parse a game")
	}
}
