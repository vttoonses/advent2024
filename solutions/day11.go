package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type Day11 struct {
	Blinks int
	stones map[string]int
}

func (data *Day11) ParseDataLine(dataLine string) error {
	data.stones = make(map[string]int)
	for _, stone := range strings.Split(strings.TrimSpace(dataLine), " ") {
		data.stones[stone] += 1
	}

	return nil
}

func (data *Day11) ProcessDataLine() error {
	return nil
}

func (data *Day11) ProcessDataSet() error {
	for ; data.Blinks > 0; data.Blinks-- {
		changes := make(map[string]int)
		for k, v := range data.stones {
			if k == "0" {
				changes["1"] += v
			} else if len(k)%2 == 1 {
				num, _ := strconv.Atoi(k)
				num *= 2024
				changes[fmt.Sprintf("%d", num)] += v
			} else {
				idx := len(k) / 2
				changes[k[:idx]] += v
				newStone := strings.TrimLeft(k[idx:], "0")
				if newStone == "" {
					newStone = "0"
				}
				changes[newStone] += v
			}
		}
		data.stones = changes
	}
	return nil
}

func (data *Day11) Solve() (int, error) {
	total := 0
	for _, v := range data.stones {
		total += v
	}
	return total, nil
}
