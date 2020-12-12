package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

// Day 5 of Advent of Code challenge
// Link: https://adventofcode.com/2020/day/5

func processSeatIDs(lines []string) []int {
	var seats []int

	for _, line := range lines {
		start, end := 0, 127
		row, col := 0, 0

		for i := 0; i < 7; i++ {
			mid := (start + end) / 2

			if line[i] == 'F' {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}

		row = start
		start, end = 0, 7

		for i := 7; i < 10; i++ {
			mid := (start + end) / 2

			if line[i] == 'L' {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}

		col = start
		seats = append(seats, row*8+col)
	}

	return seats
}

func solution1() int {
	var seats []int = processSeatIDs(parseInput())
	return getHighestSeatID(seats)
}

func solution2() int {
	var seats []int = processSeatIDs(parseInput())
	var seatMap map[int]bool = make(map[int]bool)

	for _, seatID := range seats {
		seatMap[seatID] = true
	}

	for _, seatID := range seats {
		if !seatMap[seatID+1] && seatMap[seatID+2] {
			return seatID + 1
		}
	}

	return -1
}

func getHighestSeatID(seats []int) int {
	var max int = math.MinInt32
	for _, val := range seats {
		if val > max {
			max = val
		}
	}
	return max
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

func main() {
	var result1 = solution1()
	var result2 = solution2()
	fmt.Printf("Part 1: Highest seatID: %d", result1)
	fmt.Println("")
	fmt.Printf("Part 2: Your seatID: %d", result2)
}
