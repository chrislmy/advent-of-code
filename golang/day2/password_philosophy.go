package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Day 2 of Advent of Code challenge
// Link: https://adventofcode.com/2020/day/2

// PasswordPolicy to represent password policies
type PasswordPolicy struct {
	value    byte
	minimum  int64
	maximum  int64
	password string
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

func getPassWordPolicy(input string) *PasswordPolicy {
	var segments []string = strings.Split(input, ":")
	var passwordPolicySegment []string = strings.Split(segments[0], " ")
	var minMax []string = strings.Split(passwordPolicySegment[0], "-")
	min, _ := strconv.ParseInt(minMax[0], 10, 64)
	max, _ := strconv.ParseInt(minMax[1], 10, 64)

	return &PasswordPolicy{passwordPolicySegment[1][0], min, max, strings.Trim(segments[1], " ")}
}

func solution(part int) int {
	var lines []string = parseInput()
	var result int = 0

	for _, line := range lines {
		passwordPolicy := getPassWordPolicy(line)
		var isValid bool = false
		if part == 1 {
			isValid = isValidPasswordPart1(passwordPolicy)
		} else {
			isValid = isValidPasswordPart2(passwordPolicy)
		}

		if isValid {
			result++
		}
	}

	return result
}

func isValidPasswordPart1(passwordPolicy *PasswordPolicy) bool {
	var count int64 = 0
	for _, char := range passwordPolicy.password {
		if byte(char) == passwordPolicy.value {
			count++
		}
	}
	return count >= passwordPolicy.minimum && count <= passwordPolicy.maximum
}

func isValidPasswordPart2(passwordPolicy *PasswordPolicy) bool {
	var result bool = false
	if passwordPolicy.password[passwordPolicy.minimum-1] == passwordPolicy.value {
		result = true
	}
	if passwordPolicy.password[passwordPolicy.maximum-1] == passwordPolicy.value {
		result = !result
	}

	return result
}

func main() {
	var result1 = solution(1)
	var result2 = solution(2)
	fmt.Printf("Part 1: Number of valid passwords: %d", result1)
	fmt.Println("")
	fmt.Printf("Part 2: Number of valid passwords: %d", result2)
}
