package solutions

// % go run main.go 7
// 8401132154762 <nil>
// % go run main.go 7b
// 95297119227552 <nil>

import (
	"fmt"
	"strconv"
	"strings"
)

type Day07a struct {
	line    string
	valid   []int
	AddCats bool
}

type EquationPart struct {
	running int
	add     *EquationPart
	mult    *EquationPart
	cat     *EquationPart
}

func (data *Day07a) ParseDataLine(dataLine string) error {
	data.line = dataLine
	return nil
}

func (data *Day07a) ProcessDataLine() error {
	tmp := strings.Split(data.line, ":")
	target, _ := strconv.Atoi(tmp[0])
	tmp = strings.Split(strings.TrimSpace(tmp[1]), " ")
	operands := []int{}

	for _, str := range tmp {
		op, _ := strconv.Atoi(str)
		operands = append(operands, op)
	}

	if isValid(target, operands, data.AddCats) {
		data.valid = append(data.valid, target)
	}
	return nil
}

func (data *Day07a) ProcessDataSet() error {

	return nil
}

func (data *Day07a) Solve() (int, error) {
	total := 0
	for _, val := range data.valid {
		total += val
	}

	return total, nil
}

func isValid(target int, operands []int, addCats bool) bool {
	equations := EquationPart{}
	if len(operands) == 0 {
		return false
	}
	equations.running = operands[0]
	return buildTree(&equations, operands[1:], addCats, target)
}

func buildTree(equation *EquationPart, operands []int, addCats bool, target int) bool {
	if len(operands) == 0 {
		return equation.running == target
	}

	nextOp := operands[0]

	equation.add = &EquationPart{
		running: equation.running + nextOp,
	}

	equation.mult = &EquationPart{
		running: equation.running * nextOp,
	}

	if addCats {
		catNextStr := fmt.Sprintf("%d%d", equation.running, nextOp)
		catNext, _ := strconv.Atoi(catNextStr)
		equation.cat = &EquationPart{
			running: catNext,
		}
	}

	if buildTree(equation.add, operands[1:], addCats, target) {
		return true
	}

	if buildTree(equation.mult, operands[1:], addCats, target) {
		return true
	}

	if addCats && buildTree(equation.cat, operands[1:], addCats, target) {
		return true
	}

	return false
}
