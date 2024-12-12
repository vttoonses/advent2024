package solutions

import (
	"slices"
)

type Day10 struct {
	trailHeads []Position
	scores     []int
	ratings    []int
	topo       [][]int
	width      int
	height     int
	GetRatings bool
}

func (data *Day10) ParseDataLine(dataLine string) error {
	data.topo = append(data.topo, []int{})
	for _, r := range dataLine {
		data.topo[data.height] = append(data.topo[data.height], int(r-'0'))
	}
	return nil
}

func (data *Day10) ProcessDataLine() error {
	for x, r := range data.topo[data.height] {
		if r == 0 {
			data.trailHeads = append(data.trailHeads, Position{x: x, y: data.height})
		}
	}

	if data.width == 0 {
		data.width = len(data.topo[data.height])
	}
	data.height++

	return nil
}

func (data *Day10) ProcessDataSet() error {
	for _, th := range data.trailHeads {
		summits := []Position{}
		rating := 0
		score := data.getScore(th, 0, &summits, th, &rating)
		if data.GetRatings {
			data.ratings = append(data.ratings, rating)
		} else {
			data.scores = append(data.scores, score)
		}
	}
	return nil
}

func (data *Day10) Solve() (int, error) {
	total := 0
	if data.GetRatings {
		for _, rating := range data.ratings {
			total += rating
		}
	} else {
		for _, score := range data.scores {
			total += score
		}
	}
	return total, nil
}

func (data *Day10) getScore(current Position, targetElevation int, summits *[]Position, previous Position, rating *int) int {
	if current == previous && targetElevation != 0 {
		return 0
	}

	if current.x < 0 || current.y < 0 || current.x >= data.width || current.y >= data.height {
		return 0
	}

	if data.topo[current.y][current.x] == targetElevation {
		if targetElevation == 9 {
			*rating = *rating + 1
			if !slices.Contains(*summits, current) {
				*summits = append(*summits, current)
				return 1
			}
			return 0
		}

		scores := 0
		nextElevation := targetElevation + 1
		scores += data.getScore(Position{x: current.x, y: current.y - 1}, nextElevation, summits, current, rating)
		scores += data.getScore(Position{x: current.x + 1, y: current.y}, nextElevation, summits, current, rating)
		scores += data.getScore(Position{x: current.x, y: current.y + 1}, nextElevation, summits, current, rating)
		scores += data.getScore(Position{x: current.x - 1, y: current.y}, nextElevation, summits, current, rating)
		return scores
	}

	return 0
}
