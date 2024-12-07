package solutions

import (
	"fmt"
	"slices"
)

type Addr struct {
	x int
	y int
}

func (a Addr) String() string {
	return fmt.Sprintf("(%d,%d)", a.x, a.y)
}

type GuardDeets struct {
	loc  Addr
	face rune
}

func (g GuardDeets) String() string {
	return fmt.Sprintf("(%d,%d,%c)", g.loc.x, g.loc.y, g.face)
}

type Day06a struct {
	Room     []string
	Guard    GuardDeets
	Start    GuardDeets
	Steps    map[string]string
	SkipChk  bool
	Loops    map[string]bool
	GetLoops bool
}

func (data *Day06a) ParseDataLine(dataLine string) error {
	data.Room = append(data.Room, dataLine)
	return nil
}

func (data *Day06a) ProcessDataLine() error {
	if data.Guard.face == 0 {
		sprites := []rune{'^', '>', 'v', '<'}
		y := len(data.Room) - 1

		for x, chr := range data.Room[y] {
			if slices.Contains(sprites, chr) {
				data.Guard.loc.x = x
				data.Guard.loc.y = y
				data.Guard.face = chr

				data.Start = data.Guard
			}
		}
	}
	return nil
}

func (data *Day06a) ProcessDataSet() error {
	return nil
}

func (data *Day06a) Solve() (int, error) {
	data.Steps = make(map[string]string)
	data.Loops = make(map[string]bool)
	data.Steps[data.Guard.loc.String()] = data.Guard.String()
	data.steps()
	if data.GetLoops {
		for loc := range data.Steps {
			data.checkForLoop(loc)
		}
		return len(data.Loops), nil
	}
	return len(data.Steps), nil
}

func (data *Day06a) steps() bool {
	for {
		nextX, nextY := next(data.Guard)
		if done(data.Room, nextX, nextY) {
			return false
		}

		if blocked(data.Room, nextX, nextY) {
			data.Guard.face = turnRight(data.Guard.face)
		} else {
			data.Guard.loc.x = nextX
			data.Guard.loc.y = nextY
			if data.Steps[data.Guard.loc.String()] == data.Guard.String() {
				return true
			}
			data.Steps[data.Guard.loc.String()] = data.Guard.String()
		}
	}
}

func (data *Day06a) checkForLoop(loc string) {
	var x, y int

	fmt.Sscanf(loc, "(%d,%d)", &x, &y)
	if x == data.Start.loc.x && y == data.Start.loc.y {
		return
	}

	rightPath := Day06a{
		Room:  make([]string, len(data.Room)),
		Guard: data.Start,
		Steps: make(map[string]string),
	}

	copy(rightPath.Room, data.Room)
	runeSlice := []rune(rightPath.Room[y])
	runeSlice[x] = '#'
	rightPath.Room[y] = string(runeSlice)

	if rightPath.steps() {
		data.Loops[loc] = true
	}
}

func next(position GuardDeets) (int, int) {
	switch position.face {
	case '^':
		return position.loc.x, position.loc.y - 1
	case '>':
		return position.loc.x + 1, position.loc.y
	case 'v':
		return position.loc.x, position.loc.y + 1
	default:
		return position.loc.x - 1, position.loc.y
	}
}

func blocked(room []string, x, y int) bool {
	return room[y][x] == '#'
}

func done(room []string, x, y int) bool {
	return x < 0 ||
		y < 0 ||
		y == len(room) ||
		x == len(room[y])
}

func turnRight(face rune) rune {
	switch face {
	case '^':
		return '>'
	case '>':
		return 'v'
	case 'v':
		return '<'
	default:
		return '^'
	}
}
