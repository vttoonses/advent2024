package solutions

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day07a struct {
	line    string
	valid   []int
	AddCats bool
}

type EquationPart struct {
	operand int
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
	equations.operand = operands[0]
	buildTree(&equations, operands[1:], addCats)
	totals := []int{}

	evalEquations(equations.operand, "+", equations.add, &totals, addCats)
	evalEquations(equations.operand, "*", equations.mult, &totals, addCats)
	if addCats {
		evalEquations(equations.operand, "||", equations.cat, &totals, addCats)
	}

	return slices.Contains(totals, target)
}

func buildTree(equation *EquationPart, operands []int, addCats bool) {
	if len(operands) == 0 {
		return
	}

	equation.add = &EquationPart{
		operand: operands[0],
	}

	equation.mult = &EquationPart{
		operand: operands[0],
	}

	if addCats {
		equation.cat = &EquationPart{
			operand: operands[0],
		}
	}

	if len(operands) > 1 {
		buildTree(equation.add, operands[1:], addCats)
		buildTree(equation.mult, operands[1:], addCats)
		if addCats {
			buildTree(equation.cat, operands[1:], addCats)
		}
	}
}

func evalEquations(current int, operator string, equation *EquationPart, totals *[]int, addCats bool) {
	if operator == "+" {
		current += equation.operand
	} else if operator == "*" {
		current *= equation.operand
	} else if addCats {
		tmp := fmt.Sprintf("%d%d", current, equation.operand)
		current, _ = strconv.Atoi(tmp)
	}

	if equation.add == nil {
		*totals = append(*totals, current)
		return
	}

	evalEquations(current, "+", equation.add, totals, addCats)
	evalEquations(current, "*", equation.mult, totals, addCats)
	if addCats {
		evalEquations(current, "||", equation.cat, totals, addCats)
	}
}
