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
