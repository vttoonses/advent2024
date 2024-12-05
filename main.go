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
	case "2":
		sol = &solutions.Day02a{Dampen: false}
		fileName = "./data/day02.txt"
	case "2b":
		sol = &solutions.Day02a{Dampen: true}
		fileName = "./data/day02.txt"
	case "3":
		sol = &solutions.Day03a{Conditionals: false}
		fileName = "./data/day03.txt"
	case "3b":
		sol = &solutions.Day03a{Conditionals: true}
		fileName = "./data/day03.txt"
	}

	shared.ReadWholeFile(fileName, sol)

	fmt.Println(sol.Solve())
}
