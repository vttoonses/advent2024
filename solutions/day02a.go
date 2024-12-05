package solutions

import (
	"sort"
	"strconv"
	"strings"
)

type Day02a struct {
	CurrentLevels []int
	Count         int
	Dampen        bool
}

func (data *Day02a) ParseDataLine(dataLine string) error {
	var levels []int

	nums := strings.Split(dataLine, " ")
	for _, num := range nums {
		converted, err := strconv.Atoi(num)
		if err != nil {
			return err
		}

		levels = append(levels, converted)
	}

	data.CurrentLevels = levels
	return nil
}

func (data *Day02a) ProcessDataLine() error {
	if processLevels(data.CurrentLevels) {
		data.Count += 1
		return nil
	}

	if data.Dampen {
		for lcv := 0; lcv < len(data.CurrentLevels); lcv++ {
			sublevels := make([]int, len(data.CurrentLevels))
			_ = copy(sublevels, data.CurrentLevels)

			sublevels = append(sublevels[:lcv], sublevels[lcv+1:]...)
			if processLevels(sublevels) {
				data.Count += 1
				return nil
			}
		}
	}

	return nil
}

func (data *Day02a) ProcessDataSet() error {

	return nil
}

func (data *Day02a) Solve() (int, error) {

	return data.Count, nil
}

func processLevels(levels []int) bool {
	asc := make([]int, len(levels))
	dec := make([]int, len(levels))

	_ = copy(asc, levels)
	_ = copy(dec, levels)

	sort.Slice(asc, func(i, j int) bool {
		return asc[i] < asc[j]
	})

	sort.Slice(dec, func(i, j int) bool {
		return dec[i] > dec[j]
	})

	if !areEqual(levels, asc) && !areEqual(levels, dec) {
		return false
	}

	if !areClose(levels) {
		return false
	}

	return true
}

func areEqual(s1 []int, s2 []int) bool {
	for idx, e := range s1 {
		if e != s2[idx] {
			return false
		}
	}

	return true
}

func areClose(levels []int) bool {
	last := len(levels) - 1

	if last == 0 {
		return true
	}

	for lcv := 0; lcv < last; lcv++ {
		diff := abs(levels[lcv+1] - levels[lcv])
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}
