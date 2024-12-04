package solutions

import (
	"fmt"
	"sort"
)

type Day01a struct {
	ListOne []int
	ListTwo []int
}

func (data *Day01a) ParseDataLine(dataLine string) error {
	var num1, num2 int

	fmt.Sscanf(dataLine, "%d %d", &num1, &num2)
	data.ListOne = append(data.ListOne, num1)
	data.ListTwo = append(data.ListTwo, num2)
	return nil
}

func (data *Day01a) ProcessDataLine() error {

	return nil
}

func (data *Day01a) ProcessDataSet() error {
	sort.Slice(data.ListOne, func(i, j int) bool {
		return data.ListOne[i] < data.ListOne[j]
	})
	sort.Slice(data.ListTwo, func(i, j int) bool {
		return data.ListTwo[i] < data.ListTwo[j]
	})

	return nil
}

func (data *Day01a) Solve() (int, error) {
	total := 0

	for lcv := 0; lcv < len(data.ListOne); lcv++ {
		if data.ListOne[lcv] < data.ListTwo[lcv] {
			total += (data.ListTwo[lcv] - data.ListOne[lcv])
		} else {
			total += (data.ListOne[lcv] - data.ListTwo[lcv])
		}
	}
	return total, nil
}
