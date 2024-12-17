package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aoc2024/shared"
	"github.com/aoc2024/solutions"
)

func main() {
	var sol solutions.Solution
	var fileName string

	startTime := time.Now()

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
	case "4":
		sol = &solutions.Day04a{}
		fileName = "./data/day04.txt"
	case "4b":
		sol = &solutions.Day04b{}
		fileName = "./data/day04.txt"
	case "5":
		sol = &solutions.Day05a{ProcessingUpdates: false, Fix: false}
		fileName = "./data/day05.txt"
	case "5b":
		sol = &solutions.Day05a{ProcessingUpdates: false, Fix: true}
		fileName = "./data/day05.txt"
	case "6":
		sol = &solutions.Day06a{}
		fileName = "./data/day06.txt"
	case "6b":
		sol = &solutions.Day06a{GetLoops: true}
		fileName = "./data/day06.txt"
	case "7":
		sol = &solutions.Day07a{}
		fileName = "./data/day07.txt"
	case "7b":
		sol = &solutions.Day07a{AddCats: true}
		fileName = "./data/day07.txt"
	case "8":
		sol = &solutions.Day08a{}
		fileName = "./data/day08.txt"
	case "8b":
		sol = &solutions.Day08a{Harmonics: true}
		fileName = "./data/day08.txt"
	case "9":
		sol = &solutions.Day09a{}
		fileName = "./data/day09.txt"
	case "9b":
		sol = &solutions.Day09a{ByFile: true}
		fileName = "./data/day09.txt"
	case "10":
		sol = &solutions.Day10{}
		fileName = "./data/day10.txt"
	case "10b":
		sol = &solutions.Day10{GetRatings: true}
		fileName = "./data/day10.txt"
	case "11":
		sol = &solutions.Day11{Blinks: 25}
		fileName = "./data/day11.txt"
	case "11b":
		sol = &solutions.Day11{Blinks: 75}
		fileName = "./data/day11.txt"
	case "12":
		sol = &solutions.Day12{}
		fileName = "./data/day12.txt"
	case "12b":
		sol = &solutions.Day12{Bulk: true}
		fileName = "./data/day12.txt"
	case "13":
		sol = &solutions.Day13{}
		fileName = "./data/day13.txt"
	case "13b":
		sol = &solutions.Day13{TenQ: true}
		fileName = "./data/day13.txt"
	case "14":
		sol = &solutions.Day14{Width: 101, Height: 103, Seconds: 100}
		fileName = "./data/day14.txt"
	case "14b":
		sol = &solutions.Day14{Width: 101, Height: 103, Seconds: 100, Tree: true}
		fileName = "./data/day14.txt"
	case "15":
		sol = &solutions.Day15{}
		fileName = "./data/day15.txt"
	case "15b":
		sol = &solutions.Day15{Double: true}
		fileName = "./data/day15.txt"
	case "16":
		sol = &solutions.Day16{}
		fileName = "./data/day16.txt"
	case "16b":
		sol = &solutions.Day16{Tiles: true}
		fileName = "./data/day16.txt"
	case "t":
		sol = &solutions.Day16{Tiles: true}
		fileName = "./data/test.txt"
	}

	shared.ReadWholeFile(fileName, sol)

	fmt.Println(sol.Solve())
	endTime := time.Now()
	diff := endTime.Sub(startTime).Milliseconds()

	fmt.Printf("Time difference: %d milliseconds\n", diff)
}
