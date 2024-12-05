package solutions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Day05a struct {
	Rules             map[int][]int
	Updates           []int
	Middles           []int
	ProcessingUpdates bool
	Fix               bool
}

func (data *Day05a) ParseDataLine(dataLine string) error {
	if dataLine == "" {
		data.ProcessingUpdates = true
		return nil
	}

	if !data.ProcessingUpdates {
		var page1, page2 int

		if data.Rules == nil {
			data.Rules = make(map[int][]int)
		}

		fmt.Sscanf(dataLine, "%d|%d", &page1, &page2)
		data.Rules[page1] = append(data.Rules[page1], page2)
	} else {
		data.Updates = nil
		for _, numStr := range strings.Split(dataLine, ",") {
			num, _ := strconv.Atoi(numStr)
			data.Updates = append(data.Updates, num)
		}
	}
	return nil
}

func (data *Day05a) ProcessDataLine() error {
	if data.ProcessingUpdates && data.Updates != nil {
		if isGoodOrder(data.Updates, data.Rules) {
			if !data.Fix {
				middleVal := data.Updates[len(data.Updates)/2]
				data.Middles = append(data.Middles, middleVal)
			}
		} else if data.Fix {
			sort.Slice(data.Updates, func(i, j int) bool {
				return mustPreceed(data.Rules[data.Updates[i]], data.Updates[j])
			})

			middleVal := data.Updates[len(data.Updates)/2]
			data.Middles = append(data.Middles, middleVal)
		}
	}
	return nil
}

func (data *Day05a) ProcessDataSet() error {

	return nil
}

func (data *Day05a) Solve() (int, error) {
	total := 0

	for _, num := range data.Middles {
		total += num
	}

	return total, nil
}

func mustPreceed(rule []int, target int) bool {
	for _, num := range rule {
		if num == target {
			return true
		}
	}

	return false
}

func isGoodOrder(order []int, rules map[int][]int) bool {
	for idx, num := range order {
		if idx != len(order)-1 {
			for _, ruleNum := range order[idx+1:] {
				if mustPreceed(rules[ruleNum], num) {
					return false
				}
			}
		}
	}

	return true
}
