package shared

import (
	"bufio"
	"fmt"
	"os"

	"github.com/aoc2024/solutions"
)

func ReadWholeFile(fileName string, sol solutions.Solution) error {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sol.ParseDataLine(scanner.Text())
		sol.ProcessDataLine()
	}

	if err := scanner.Err(); err == nil {
		sol.ProcessDataSet()
		return nil
	} else {
		return scanner.Err()
	}
}
