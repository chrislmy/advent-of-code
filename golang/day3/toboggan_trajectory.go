package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Day 3 of Advent of Code challenge
// Link: https://adventofcode.com/2020/day/3

// Slope to represent slope values for part 2 solution
type Slope struct {
	x int
	y int
}

func parseInput() []string {
	var grid []string
	inputFile, error := os.Open("input.txt")
	if error != nil {
		log.Fatal(error)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if error := scanner.Err(); error != nil {
		log.Fatal(error)
	}

	return grid
}

func solution1() int {
	var grid []string = parseInput()
	return countTrees(3, 1, grid)
}

func solution2() int {
	var slopes [5]Slope = [5]Slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	var grid []string = parseInput()
	var result int = 1

	for _, slope := range slopes {
		result *= countTrees(slope.x, slope.y, grid)
	}

	return result
}

func countTrees(slopeX, slopeY int, grid []string) int {
	var row int = 0
	var col int = 0
	var numTrees int = 0

	for row < len(grid) {
		if grid[row][col] == '#' {
			numTrees++
		}
		col = (col + slopeX) % len(grid[row])
		row = row + slopeY
	}

	return numTrees
}

func main() {
	var result1 = solution1()
	var result2 = solution2()
	fmt.Printf("Part 1: Number of trees: %d", result1)
	fmt.Println("")
	fmt.Printf("Part 2: Number of trees: %d", result2)
}
