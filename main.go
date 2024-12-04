package main

import (
	"fmt"
	"os"

	"github.com/aoc2024/shared"
	"github.com/aoc2024/solutions"
)

func main() {
	var sol solutions.Solution
	var fileName string

	switch os.Args[1] {
	case "1":
		sol = &solutions.Day01a{}
		fileName = "./data/day01.txt"
	case "1b":
		sol = &solutions.Day01b{}
		fileName = "./data/day01.txt"
	}

	shared.ReadWholeFile(fileName, sol)

	fmt.Println(sol.Solve())
}
