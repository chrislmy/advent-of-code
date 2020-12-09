package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// Day 1 of Advent of Code challenge
// Link: https://adventofcode.com/2020/day/1

var TARGET int64 = 2020

func parseInput() []int64 {
	var result []int64
	inputFile, error := os.Open("input.txt")
	if error != nil {
		log.Fatal(error)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		result = append(result, num)
	}

	if error := scanner.Err(); error != nil {
		log.Fatal(error)
	}

	return result
}

func solutionPart1() [2]int64 {
	var input []int64 = parseInput()
	var set map[int64]bool = make(map[int64]bool)

	for _, num := range input {
		var remainder int64 = TARGET - num
		if set[remainder] {
			return [2]int64{num, remainder}
		}

		set[num] = true
	}

	return [2]int64{-1, -1}
}

func solutionPart2() [3]int64 {
	var input []int64 = parseInput()
	sort.Slice(input, func(e1, e2 int) bool {
		return input[e1] < input[e2]
	})

	for i := 0; i < len(input)-2; i++ {
		start := i + 1
		end := len(input) - 1

		for start < end {
			sum := input[i] + input[start] + input[end]

			if sum == TARGET {
				return [3]int64{input[i], input[start], input[end]}
			} else if sum < TARGET {
				start++
			} else {
				end--
			}
		}
	}

	return [3]int64{-1, -1, -1}
}

func main() {
	var result1 [2]int64 = solutionPart1()
	var result2 [3]int64 = solutionPart2()
	fmt.Printf("Pair: [%d, %d], Result: %d", result1[0], result1[1], result1[0]*result1[1])
	fmt.Println("")
	fmt.Printf("Triplet: [%d, %d, %d], Result: %d", result2[0], result2[1], result2[2], result2[0]*result2[1]*result2[2])
}
