package solutions

import (
	"fmt"
	"math"
	"slices"
)

const N = 1
const E = 2
const S = 3
const W = 4

type node struct {
	face     int
	location Position
	score    int
	path     []Position
}

type costs struct {
	cost int
	face int
}

type Day16 struct {
	maze        [][]rune
	costs       [][]costs
	start       Position
	end         Position
	score       int
	path        []Position
	placesToSit map[int][][]Position
	Tiles       bool
}

func (data *Day16) ParseDataLine(dataLine string) error {
	data.maze = append(data.maze, []rune(dataLine))
	data.costs = append(data.costs, make([]costs, len(dataLine)))
	if col := slices.Index(data.maze[len(data.maze)-1], 'S'); col >= 0 {
		data.start = Position{col, len(data.maze) - 1}
	}
	if col := slices.Index(data.maze[len(data.maze)-1], 'E'); col >= 0 {
		data.end = Position{col, len(data.maze) - 1}
	}
	return nil
}

func (data *Day16) ProcessDataLine() error {
	return nil
}

func (data *Day16) ProcessDataSet() error {
	// data.placesToSit = make(map[int][][]Position)
	// start := node{face: E, location: data.start, score: 0}
	// start.path = append(start.path, data.start)
	// if data.maze[start.location.y-1][start.location.x] == '.' {
	// 	data.traceRoutes(start, N, 1001)
	// }
	// if data.maze[start.location.y+1][start.location.x] == '.' {
	// 	data.traceRoutes(start, S, 1001)
	// }
	// if data.maze[start.location.y][start.location.x-1] == '.' {
	// 	data.traceRoutes(start, W, 2001)
	// }
	// if data.maze[start.location.y][start.location.x+1] == '.' {
	// 	data.traceRoutes(start, E, 1)
	// }
	fmt.Println(data.aStarSearch())
	return nil
}

func (data *Day16) Solve() (int, error) {
	cost := 0
	face := E
	for lcv := 0; lcv < len(data.path)-1; lcv++ {
		first := data.path[lcv]
		second := data.path[lcv+1]
		stepCost, nextFace := moveCost(first, second, face)
		cost += stepCost
		face = nextFace
	}
	// fmt.Println("Best cost", data.score)
	// if data.Tiles {
	// 	var superSet []Position

	// 	for _, singleSet := range data.placesToSit[data.score] {
	// 		for _, tile := range singleSet {
	// 			if slices.Index(superSet, tile) < 0 {
	// 				superSet = append(superSet, tile)
	// 			}
	// 		}
	// 	}
	// 	return len(superSet) + 1, nil
	// } else {
	// 	return data.score, nil
	// }
	return cost, nil
}

// func (data *Day16) belowMin(score int) bool {
// 	if data.score == 0 {
// 		return true
// 	}

// 	return score <= data.score
// }

// func (data *Day16) traceRoutes(current node, direction int, score int) {
// 	nextLoc := nextLocation(current.location, direction)
// 	if slices.Index(current.path, nextLoc) >= 0 {
// 		return
// 	}

// 	if !data.belowMin(current.score + score) {
// 		return
// 	}
// 	if data.costs[nextLoc.y][nextLoc.x].cost >= current.score+score || data.costs[nextLoc.y][nextLoc.x].cost == 0 {
// 		data.costs[nextLoc.y][nextLoc.x].cost = current.score + score
// 	} else {
// 		return
// 	}

// 	if nextLoc == data.end {
// 		finalScore := current.score + score
// 		fmt.Println("Found End", finalScore, len(current.path))
// 		data.placesToSit[finalScore] = append(data.placesToSit[finalScore], current.path)
// 		data.score = current.score + score
// 		return
// 	}

// 	nextNode := node{face: direction, location: nextLoc, score: current.score + score}
// 	nextNode.path = append(current.path, nextLoc)
// 	left, right := getFacings(direction)

// 	if pos := nextLocation(nextLoc, direction); data.maze[pos.y][pos.x] == '.' || pos == data.end {
// 		data.traceRoutes(nextNode, direction, 1)
// 	}
// 	if pos := nextLocation(nextLoc, left); data.maze[pos.y][pos.x] == '.' || pos == data.end {
// 		data.traceRoutes(nextNode, left, 1001)
// 	}
// 	if pos := nextLocation(nextLoc, right); data.maze[pos.y][pos.x] == '.' || pos == data.end {
// 		data.traceRoutes(nextNode, right, 1001)
// 	}
// }

// func getFacings(face int) (left, right int) {
// 	left = face - 1
// 	if left < N {
// 		left = W
// 	}

// 	right = face + 1
// 	if right > W {
// 		right = N
// 	}

// 	return
// }

//	func nextLocation(location Position, face int) Position {
//		switch face {
//		case N:
//			return Position{location.x, location.y - 1}
//		case S:
//			return Position{location.x, location.y + 1}
//		case E:
//			return Position{location.x + 1, location.y}
//		default:
//			return Position{location.x - 1, location.y}
//		}
//	}
func moveCost(first, second Position, face int) (int, int) {
	cost := 0
	newFace := 0

	if second.x > first.x {
		newFace = E
		if face == newFace {
			cost += 1
		} else {
			cost += 1001
		}
	} else if second.x < first.x {
		newFace = W
		if face == newFace {
			cost += 1
		} else {
			cost += 1001
		}
	} else if second.y < first.y {
		newFace = N
		if face == newFace {
			cost += 1
		} else {
			cost += 1001
		}
	} else if second.y > first.y {
		newFace = S
		if face == newFace {
			cost += 1
		} else {
			cost += 1001
		}
	}

	return cost, newFace
}

type d16Coords struct {
	x int
	y int
}

type d16pPair struct {
	f      float64
	coords d16Coords
}

type d16Cell struct {
	parent_x int
	parent_y int
	f        float64
	g        float64
	h        float64
}

func (data *Day16) isValid(x, y int) bool {
	return y < len(data.maze) && x < len(data.maze[0])
}

func (data *Day16) isUnblocked(x, y int) bool {
	return data.maze[y][x] == '.'
}

func (data *Day16) isDestination(x, y int) bool {
	return Position{x, y} == data.end
}

func (data *Day16) calculateHValue(x, y int) float64 {
	// Using Manhattan algorithm
	return float64(abs(x-data.end.x) + abs(y-data.end.y))
}

func (data *Day16) tracePath(cells *[][]d16Cell) {
	// var path []Position

	x := data.end.x
	y := data.end.y
	pos := Position{x, y}

	for pos != data.start {
		data.path = append(data.path, pos)
		tx := (*cells)[pos.y][pos.x].parent_x
		ty := (*cells)[pos.y][pos.x].parent_y
		pos = Position{tx, ty}
	}

	data.path = append(data.path, data.start)

	// for lcv := len(data.path) - 1; lcv >= 0; lcv-- {
	// 	fmt.Println(data.path[lcv])
	// }
}

func firstOpen(collection map[d16pPair]bool) d16pPair {
	for k := range collection {
		return k
	}

	return d16pPair{}
}

func (data *Day16) aStarSearch() bool {
	var closedList [][]bool

	rows := len(data.maze)
	cols := len(data.maze[0])

	closedList = make([][]bool, rows)
	for lcv := 0; lcv < rows; lcv++ {
		closedList[lcv] = make([]bool, cols)
	}

	var cellDetails [][]d16Cell

	cellDetails = make([][]d16Cell, rows)
	for lcv := 0; lcv < rows; lcv++ {
		cellDetails[lcv] = make([]d16Cell, cols)
		for lcv2 := 0; lcv2 < cols; lcv2++ {
			cellDetails[lcv][lcv2] = d16Cell{-1, -1, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64}
		}
	}

	cellDetails[data.start.y][data.start.x] = d16Cell{data.start.y, data.start.x, 0.0, 0.0, 0.0}

	var openList map[d16pPair]bool

	openList = make(map[d16pPair]bool)
	openList[d16pPair{0.0, d16Coords{data.start.x, data.start.y}}] = true

	for len(openList) > 0 {
		p := firstOpen(openList)
		delete(openList, p)

		x := p.coords.x
		y := p.coords.y
		closedList[y][x] = true

		// North
		if data.checkPosition(Position{x, y - 1}, Position{x, y}, &cellDetails, &openList, &closedList) ||
			// East
			data.checkPosition(Position{x + 1, y}, Position{x, y}, &cellDetails, &openList, &closedList) ||
			// South
			data.checkPosition(Position{x, y + 1}, Position{x, y}, &cellDetails, &openList, &closedList) ||
			// West
			data.checkPosition(Position{x - 1, y}, Position{x, y}, &cellDetails, &openList, &closedList) {
			return true
		}
	}
	return false
}

func (data *Day16) checkPosition(pos Position, orig Position, cells *[][]d16Cell, openList *map[d16pPair]bool, closedList *[][]bool) bool {
	var gNew, hNew, fNew float64

	if data.isValid(pos.x, pos.y) {
		if data.isDestination(pos.x, pos.y) {
			(*cells)[pos.y][pos.x].parent_x = orig.x
			(*cells)[pos.y][pos.x].parent_y = orig.y
			data.tracePath(cells)
			return true
		} else if !(*closedList)[pos.y][pos.x] && data.isUnblocked(pos.x, pos.y) {
			gNew = (*cells)[orig.y][orig.x].g + 1.0
			hNew = data.calculateHValue(pos.x, pos.y)
			fNew = gNew + hNew

			if (*cells)[pos.y][pos.x].f > fNew {
				(*openList)[d16pPair{fNew, d16Coords{pos.x, pos.y}}] = true
				(*cells)[pos.y][pos.x] = d16Cell{orig.x, orig.y, fNew, gNew, hNew}
			}
		}
	}
	return false
}
