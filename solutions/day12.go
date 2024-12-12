package solutions

import (
	"slices"
)

type region struct {
	perimeter int
	positions []Position
	verticies int
}

type Day12 struct {
	plots     []string
	regions   []region
	positions []Position
	width     int
	height    int
	Bulk      bool
}

func (data *Day12) ParseDataLine(dataLine string) error {
	data.plots = append(data.plots, dataLine)
	return nil
}

func (data *Day12) ProcessDataLine() error {
	data.height = len(data.plots)
	y := data.height - 1
	if data.width == 0 {
		data.width = len(data.plots[y])
	}
	for x := range data.plots[y] {
		data.positions = append(data.positions, Position{x, y})
	}
	return nil
}

func (data *Day12) ProcessDataSet() error {
	for len(data.positions) > 0 {
		position := data.positions[0]
		data.regions = append(data.regions, region{})
		data.defineRegion(position, rune(data.plots[position.y][position.x]))
		data.countVertices(&data.regions[len(data.regions)-1])
	}
	return nil
}

func (data *Day12) Solve() (int, error) {
	total := 0
	for _, regn := range data.regions {
		if data.Bulk {
			total += (len(regn.positions) * regn.verticies)
		} else {
			total += (len(regn.positions) * regn.perimeter)
		}
	}
	return total, nil
}

func (data *Day12) defineRegion(pos Position, plant rune) {
	if pos.x < 0 || pos.y < 0 || pos.x >= data.width || pos.y >= data.height {
		data.regions[len(data.regions)-1].perimeter += 1
		return
	}

	if rune(data.plots[pos.y][pos.x]) != plant {
		data.regions[len(data.regions)-1].perimeter += 1
		return
	}

	idx := slices.Index(data.positions, pos)
	if idx == -1 {
		return
	} else {
		data.regions[len(data.regions)-1].positions = append(data.regions[len(data.regions)-1].positions, pos)
		data.positions = slices.Delete(data.positions, idx, idx+1)
	}

	data.defineRegion(Position{pos.x, pos.y - 1}, plant)
	data.defineRegion(Position{pos.x + 1, pos.y}, plant)
	data.defineRegion(Position{pos.x, pos.y + 1}, plant)
	data.defineRegion(Position{pos.x - 1, pos.y}, plant)
}

func (data *Day12) countVertices(regn *region) {
	for _, pos := range regn.positions {
		if data.outerUpperLeft(pos) {
			regn.verticies += 1
		}
		if data.outerUpperRight(pos) {
			regn.verticies += 1
		}
		if data.outerLowerLeft(pos) {
			regn.verticies += 1
		}
		if data.outerLowerRight(pos) {
			regn.verticies += 1
		}
		if data.innerUpperLeft(pos) {
			regn.verticies += 1
		}
		if data.innerUpperRight(pos) {
			regn.verticies += 1
		}
		if data.innerLowerLeft(pos) {
			regn.verticies += 1
		}
		if data.innerLowerRight(pos) {
			regn.verticies += 1
		}
	}
}

func (data *Day12) outerUpperLeft(pos Position) bool {
	if pos.x-1 < 0 && pos.y-1 < 0 {
		return true
	}

	plant := data.plots[pos.y][pos.x]

	if pos.x-1 < 0 && data.plots[pos.y-1][pos.x] != plant {
		return true
	}

	if pos.x-1 < 0 {
		return false
	}

	if pos.y-1 < 0 && data.plots[pos.y][pos.x-1] != plant {
		return true
	}

	return data.plots[pos.y][pos.x-1] != plant &&
		data.plots[pos.y-1][pos.x] != plant
}

func (data *Day12) outerUpperRight(pos Position) bool {
	if pos.x+1 >= data.width && pos.y-1 < 0 {
		return true
	}

	plant := data.plots[pos.y][pos.x]

	if pos.x+1 >= data.width && data.plots[pos.y-1][pos.x] != plant {
		return true
	}

	if pos.x+1 >= data.width {
		return false
	}

	if pos.y-1 < 0 && data.plots[pos.y][pos.x+1] != plant {
		return true
	}

	return data.plots[pos.y][pos.x+1] != plant &&
		data.plots[pos.y-1][pos.x] != plant
}

func (data *Day12) outerLowerLeft(pos Position) bool {
	if pos.x-1 < 0 && pos.y+1 >= data.height {
		return true
	}

	plant := data.plots[pos.y][pos.x]

	if pos.x-1 < 0 && data.plots[pos.y+1][pos.x] != plant {
		return true
	}

	if pos.x-1 < 0 {
		return false
	}

	if pos.y+1 >= data.height && data.plots[pos.y][pos.x-1] != plant {
		return true
	}

	return data.plots[pos.y][pos.x-1] != plant &&
		data.plots[pos.y+1][pos.x] != plant
}

func (data *Day12) outerLowerRight(pos Position) bool {
	if pos.x+1 >= data.width && pos.y+1 >= data.height {
		return true
	}

	plant := data.plots[pos.y][pos.x]

	if pos.x+1 >= data.width && data.plots[pos.y+1][pos.x] != plant {
		return true
	}

	if pos.x+1 >= data.width {
		return false
	}

	if pos.y+1 >= data.height && data.plots[pos.y][pos.x+1] != plant {
		return true
	}

	return data.plots[pos.y][pos.x+1] != plant &&
		data.plots[pos.y+1][pos.x] != plant
}

func (data *Day12) innerUpperLeft(pos Position) bool {
	if pos.x+1 >= data.width || pos.y+1 >= data.height {
		return false
	}

	plant := data.plots[pos.y][pos.x]

	return data.plots[pos.y][pos.x+1] == plant &&
		data.plots[pos.y+1][pos.x+1] != plant &&
		data.plots[pos.y+1][pos.x] == plant
}

func (data *Day12) innerUpperRight(pos Position) bool {
	if pos.x-1 < 0 || pos.y+1 >= data.height {
		return false
	}

	plant := data.plots[pos.y][pos.x]

	return data.plots[pos.y][pos.x-1] == plant &&
		data.plots[pos.y+1][pos.x-1] != plant &&
		data.plots[pos.y+1][pos.x] == plant
}

func (data *Day12) innerLowerLeft(pos Position) bool {
	if pos.x+1 >= data.width || pos.y-1 < 0 {
		return false
	}

	plant := data.plots[pos.y][pos.x]

	return data.plots[pos.y][pos.x+1] == plant &&
		data.plots[pos.y-1][pos.x+1] != plant &&
		data.plots[pos.y-1][pos.x] == plant
}

func (data *Day12) innerLowerRight(pos Position) bool {
	if pos.x-1 < 0 || pos.y-1 < 0 {
		return false
	}

	plant := data.plots[pos.y][pos.x]

	return data.plots[pos.y][pos.x-1] == plant &&
		data.plots[pos.y-1][pos.x-1] != plant &&
		data.plots[pos.y-1][pos.x] == plant
}
