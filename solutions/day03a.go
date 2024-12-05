package solutions

import (
	"regexp"
	"strconv"
)

type Day03a struct {
	Memory       string
	Operations   []string
	Products     []int
	Conditionals bool
	Do           bool
}

var instructions *regexp.Regexp
var operands *regexp.Regexp

func (data *Day03a) ParseDataLine(dataLine string) error {
	data.Memory += dataLine
	return nil
}

func (data *Day03a) ProcessDataLine() error {
	return nil
}

func (data *Day03a) ProcessDataSet() error {
	if instructions == nil {
		instructions, _ = regexp.Compile(`(mul\([0-9]{1,3},[0-9]{1,3}\))|(do\(\))|(don't\(\))`)
	}

	if operands == nil {
		operands, _ = regexp.Compile(`[0-9]{1,3}`)
	}

	data.Do = true

	for {
		loc := instructions.FindIndex([]byte(data.Memory))
		if loc == nil {
			break
		}

		instruction := data.Memory[loc[0]:loc[1]]
		data.Memory = data.Memory[loc[1]:]

		switch instruction {
		case "don't()":
			if data.Conditionals {
				data.Do = false
			}
		case "do()":
			data.Do = true
		default:
			if data.Do {
				one, two := parseNumbers(operands, instruction)
				data.Products = append(data.Products, one*two)
			}
		}
	}
	return nil
}

func (data *Day03a) Solve() (int, error) {
	total := 0
	for _, p := range data.Products {
		total += p
	}
	return total, nil
}

func parseNumbers(pattern *regexp.Regexp, instruction string) (numOne int, numTwo int) {
	numStrings := pattern.FindAllString(instruction, -1)
	numOne, _ = strconv.Atoi(numStrings[0])
	numTwo, _ = strconv.Atoi(numStrings[1])

	return
}
