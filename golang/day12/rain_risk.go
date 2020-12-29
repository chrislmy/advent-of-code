package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

var directions [4]int = [4]int{1, 1, -1, -1}

// Day 12 of Advent of Code challenge
// Link: https://adventofcode.com/2020/day/12

type Ship struct {
	horizontalDistance int
	verticalDistance   int
	angle              int
}

type Instruction struct {
	action rune
	value  int
}

func (ship *Ship) eval(instruction *Instruction) {
	switch instruction.action {
	case 'N':
		ship.verticalDistance += instruction.value
		break
	case 'E':
		ship.horizontalDistance += instruction.value
		break
	case 'S':
		ship.verticalDistance -= instruction.value
		break
	case 'W':
		ship.horizontalDistance -= instruction.value
		break
	case 'L':
		ship.angle = (ship.angle - instruction.value) % 360
		break
	case 'R':
		ship.angle = (ship.angle + instruction.value) % 360
		break
	case 'F':
		directionIndex := ship.angle / 90
		if ship.angle%180 == 0 {
			ship.verticalDistance += (instruction.value * directions[directionIndex])
		} else {
			ship.horizontalDistance += (instruction.value * directions[directionIndex])
		}
		break
	}

	if ship.angle < 0 {
		ship.angle = 360 - int(math.Abs(float64(ship.angle)))
	}
}

func (ship *Ship) getDistance() int {
	var absHorizontal float64 = math.Abs(float64(ship.horizontalDistance))
	var absVertical float64 = math.Abs(float64(ship.verticalDistance))
	return int(absHorizontal + absVertical)
}

func solution() int {
	var lines []string = parseInput()
	var ship *Ship = &Ship{0, 0, 90}
	for _, line := range lines {
		instruction := getInstruction(line)
		ship.eval(instruction)
	}
	return ship.getDistance()
}

func getInstruction(input string) *Instruction {
	action := input[0]
	value := 0
	for i := 1; i < len(input); i++ {
		value = int(input[i]-'0') + (value * 10)
	}
	return &Instruction{rune(action), value}
}

func parseInput() []string {
	var result []string
	inputFile, error := os.Open("input.txt")
	if error != nil {
		log.Fatal(error)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if error := scanner.Err(); error != nil {
		log.Fatal(error)
	}

	return result
}

func main() {
	result := solution()
	fmt.Printf("Part 1: Distance of ship: %d", result)
}
