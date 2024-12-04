package shared

import (
	"bufio"
	"os"

	"github.com/aoc2024/solutions"
)

func ReadWholeFile(fileName string, sol solutions.Solution) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
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
