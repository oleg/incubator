package main

import "testing"

func TestDay1p1(t *testing.T) {
	result := day1p1(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)
	if result != "142" {
		t.Errorf("got %s, expected 142", result)
	}
}

func TestDay1p2(t *testing.T) {
	result := day1p2(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`)
	if result != "281" {
		t.Errorf("got %s, expected 281", result)
	}
}

func TestDay2p1(t *testing.T) {
	result := day2p1(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`)
	if result != "8" {
		t.Errorf("got %s, expected 8", result)
	}
}

func TestDayParseGame(t *testing.T) {
	result := parseGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	g := game{
		id: 1,
		rounds: []round{
			{4, 0, 3},
			{1, 2, 6},
			{0, 2, 0},
		},
	}
	if result.id != g.id {
		t.Errorf("got %d, expected %d", result.id, g.id)
	}
	for i := range g.rounds {
		if result.rounds[i] != g.rounds[i] {
			t.Errorf("got %v, expected %v", result.rounds[i], g.rounds[i])
		}
	}
}
