package solutions

import (
	"slices"
)

type Position struct {
	x int
	y int
}

type Day08a struct {
	width     int
	height    int
	antennae  map[rune][]Position
	antinodes []Position
	line      string
	Harmonics bool
}

func (data *Day08a) ParseDataLine(dataLine string) error {
	data.line = dataLine
	return nil
}

func (data *Day08a) ProcessDataLine() error {
	if data.width == 0 {
		data.width = len(data.line)
	}
	data.height += 1

	if data.antennae == nil {
		data.antennae = make(map[rune][]Position)
	}

	for pos, chr := range data.line {
		if chr != '.' {
			data.antennae[chr] = append(data.antennae[chr], Position{
				x: pos,
				y: data.height - 1,
			})
		}
	}
	return nil
}

func (data *Day08a) ProcessDataSet() error {
	for _, positions := range data.antennae {
		for lcv, pos1 := range positions {
			for _, pos2 := range positions[lcv+1:] {
				nodes := determineAntinodes(pos1, pos2, data.width, data.height, data.Harmonics)

				for _, node := range nodes {
					if !slices.Contains(data.antinodes, node) {
						data.antinodes = append(data.antinodes, node)
					}
				}
			}
		}
	}
	return nil
}

func (data *Day08a) Solve() (int, error) {
	return len(data.antinodes), nil
}

func determineAntinodes(p1, p2 Position, width, height int, harmonics bool) (nodes []Position) {
	var xDif, yDif int

	xDif = p1.x - p2.x
	yDif = p1.y - p2.y

	nodes = append(nodes, tracePath(p1, xDif, yDif, width, height, harmonics)...)
	nodes = append(nodes, tracePath(p1, -xDif, -yDif, width, height, harmonics)...)

	return
}

func tracePath(point Position, xDif, yDif, width, height int, harmonics bool) (points []Position) {
	for {
		if point.x < 0 || point.y < 0 || point.x >= width || point.y >= height {
			return
		}

		points = append(points, point)

		if !harmonics {
			return
		}

		point.x += xDif
		point.y += yDif
	}
}
