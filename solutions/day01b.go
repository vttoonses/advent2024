package solutions

import (
	"fmt"
	"sort"
)

type Day01b struct {
	ListOne []int
	ListTwo []int
	Counts  []int
}

func (data *Day01b) ParseDataLine(dataLine string) error {
	var num1, num2 int

	fmt.Sscanf(dataLine, "%d %d", &num1, &num2)
	data.ListOne = append(data.ListOne, num1)
	data.ListTwo = append(data.ListTwo, num2)
	return nil
}

func (data *Day01b) ProcessDataLine() error {

	return nil
}

func (data *Day01b) ProcessDataSet() error {
	sort.Slice(data.ListOne, func(i, j int) bool {
		return data.ListOne[i] < data.ListOne[j]
	})
	sort.Slice(data.ListTwo, func(i, j int) bool {
		return data.ListTwo[i] < data.ListTwo[j]
	})

	for lcv := 0; lcv < len(data.ListOne); lcv++ {
		count := instanceCount(data.ListOne[lcv], data.ListTwo)
		data.Counts = append(data.Counts, count*data.ListOne[lcv])
	}

	return nil
}

func (data *Day01b) Solve() (int, error) {
	total := 0

	for lcv := 0; lcv < len(data.Counts); lcv++ {
		total += data.Counts[lcv]
	}

	return total, nil
}

func instanceCount(target int, collection []int) int {
	count := len(collection)

	if count == 0 {
		return 0
	}

	i, found := sort.Find(count, func(i int) int {
		if target == collection[i] {
			return 0
		} else if target < collection[i] {
			return -1
		} else {
			return 1
		}
	})

	if !found {
		return 0
	}

	subcollection := collection[i+1:]
	return 1 + instanceCount(target, subcollection)
}
