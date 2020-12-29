package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	operation string
	argument  int
}

func solution1() int {
	var visited map[int]bool = make(map[int]bool)
	var value int = 0
	var index int = 0
	var lines []string = parseInput()

	for index < len(lines) {
		if visited[index] {
			return value
		}

		visited[index] = true
		instruction := getInstruction(lines[index])
		if instruction.operation == "jmp" {
			index = index + instruction.argument
		} else {
			value = instruction.eval(value)
			index++
		}
	}

	return value
}

func getInstruction(input string) *Instruction {
	var segments []string = strings.Split(input, " ")
	var operation string = segments[0]
	argument, _ := strconv.Atoi(segments[1])
	return &Instruction{operation, argument}
}

func (instruction Instruction) eval(currentVal int) int {
	switch instruction.operation {
	case "nop":
		return currentVal
	case "acc":
		return currentVal + instruction.argument
	default:
		return currentVal
	}
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
	var result1 = solution1()
	fmt.Printf("Part 1: Value of accumulator: %d", result1)
}
