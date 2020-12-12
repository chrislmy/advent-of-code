package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Day 6 of Advent of Code challenge
// Link: https://adventofcode.com/2020/day/6

var DELIMITER string = ","

func solution1() int {
	var result int = 0
	var groups []string = parseInput()

	for _, group := range groups {
		result += countAnswers(group)
	}

	return result
}

func solution2() int {
	var result int = 0
	var groups []string = parseInput()

	for _, group := range groups {
		result += countIntersection(group)
	}

	return result
}

func countAnswers(group string) int {
	var answerMap map[rune]bool = make(map[rune]bool)

	for _, char := range group {
		if char != rune(DELIMITER[0]) {
			answerMap[char] = true
		}
	}

	return len(answerMap)
}

func countIntersection(group string) int {
	var answerMap map[rune]int = make(map[rune]int)
	var result int = 0
	var answers []string = strings.Split(group, DELIMITER)

	for _, answer := range answers {
		for _, char := range answer {
			answerMap[char]++
		}
	}

	for _, count := range answerMap {
		if count == len(answers) {
			result++
		}
	}

	return result
}

func parseInput() []string {
	var groups []string
	inputFile, error := os.Open("input.txt")
	if error != nil {
		log.Fatal(error)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var stringBuilder strings.Builder

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			groups = append(groups, stringBuilder.String()[0:stringBuilder.Len()-1])
			stringBuilder.Reset()
		} else {
			stringBuilder.WriteString(line)
			stringBuilder.WriteString(DELIMITER)
		}
	}

	if error := scanner.Err(); error != nil {
		log.Fatal(error)
	}

	groups = append(groups, stringBuilder.String()[0:stringBuilder.Len()-1])
	return groups
}

func main() {
	var result1 = solution1()
	var result2 = solution2()
	fmt.Printf("Part 1: Number of asnwers: %d", result1)
	fmt.Println("")
	fmt.Printf("Part 2: Number of asnwers: %d", result2)
}
