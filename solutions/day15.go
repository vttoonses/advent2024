package solutions

import (
	"fmt"
	"slices"
)

type Day15 struct {
	Double      bool
	robot       Position
	warehouse   [][]rune
	commands    []rune
	getCommands bool
	y           int
	w           int
}

func (data *Day15) ParseDataLine(dataLine string) error {
	runes := []rune(dataLine)
	if data.getCommands {
		data.commands = append(data.commands, runes...)
	} else if dataLine == "" {
		data.getCommands = true
	} else {
		data.w = len(runes)
		if data.Double {
			data.w *= 2
		}
		row := make([]rune, data.w)
		data.warehouse = append(data.warehouse, row)
		for x, r := range runes {
			pos := Position{x, data.y}
			if data.Double {
				pos.x *= 2
			}

			switch r {
			case '@':
				data.robot = pos
				data.warehouse[pos.y][pos.x] = '@'
				if data.Double {
					data.warehouse[pos.y][pos.x+1] = '.'
				}
			case '#':
				data.warehouse[pos.y][pos.x] = '#'
				if data.Double {
					data.warehouse[pos.y][pos.x+1] = '#'
				}
			case 'O':
				if data.Double {
					data.warehouse[pos.y][pos.x] = '['
					data.warehouse[pos.y][pos.x+1] = ']'
				} else {
					data.warehouse[pos.y][pos.x] = 'O'
				}
			default:
				data.warehouse[pos.y][pos.x] = '.'
				if data.Double {
					data.warehouse[pos.y][pos.x+1] = '.'
				}
			}
		}
		data.y += 1
	}
	return nil
}

func (data *Day15) ProcessDataLine() error {

	return nil
}

func (data *Day15) ProcessDataSet() error {
	for _, command := range data.commands {
		data.move(data.robot, command)
	}

	return nil
}

func (data *Day15) Solve() (int, error) {
	total := 0
	target := 'O'
	if data.Double {
		target = '['
	}

	for y, row := range data.warehouse {
		for x, col := range row {
			if col == target {
				total += 100*y + x
			}
		}
	}

	return total, nil
}

func (data *Day15) findAllBoxes(pos Position, direction rune) []Position {
	moveFunc := north
	switch direction {
	case '>':
		moveFunc = east
	case 'v':
		moveFunc = south
	case '<':
		moveFunc = west
	}

	r := data.warehouse[pos.y][pos.x]
	if r == '#' || r == '.' {
		return []Position{}
	}

	next := moveFunc(pos)
	result := append([]Position{pos}, data.findAllBoxes(next, direction)...)
	if r == '[' && (direction == '^' || direction == 'v') {
		result = append(result, Position{pos.x + 1, pos.y})
		result = append(result, data.findAllBoxes(Position{next.x + 1, next.y}, direction)...)
	}
	if r == ']' && (direction == '^' || direction == 'v') {
		result = append(result, Position{pos.x - 1, pos.y})
		result = append(result, data.findAllBoxes(Position{next.x - 1, next.y}, direction)...)
	}

	return result
}

func (data *Day15) move(pos Position, direction rune) bool {
	moveFunc := north
	switch direction {
	case '>':
		moveFunc = east
	case 'v':
		moveFunc = south
	case '<':
		moveFunc = west
	}

	boxes := data.findAllBoxes(pos, direction)
	for len(boxes) > 0 {
		topEdge := make(map[Position]int)

		for lcv, box := range boxes {
			next := moveFunc(box)
			nextR := data.warehouse[next.y][next.x]
			if nextR == '#' {
				return false
			}
			if nextR == '.' {
				topEdge[box] = lcv
			}
		}

		for top := range topEdge {
			for lcv := slices.Index(boxes, top); lcv >= 0; lcv = slices.Index(boxes, top) {
				boxes = slices.Delete(boxes, lcv, lcv+1)
			}
			next := moveFunc(top)
			data.warehouse[next.y][next.x] = data.warehouse[top.y][top.x]
			data.warehouse[top.y][top.x] = '.'
			if data.warehouse[next.y][next.x] == '@' {
				data.robot = next
			}
		}
	}
	return true
}

func north(pos Position) Position {
	return Position{pos.x, pos.y - 1}
}

func east(pos Position) Position {
	return Position{pos.x + 1, pos.y}
}

func south(pos Position) Position {
	return Position{pos.x, pos.y + 1}
}

func west(pos Position) Position {
	return Position{pos.x - 1, pos.y}
}

func (data *Day15) displayWH() {
	for _, row := range data.warehouse {
		for _, run := range row {
			fmt.Printf("%c", run)
		}
		fmt.Println()
	}
}
