package solutions

import (
	"fmt"
	"strings"
)

type clawMachine struct {
	a     Position
	b     Position
	prize Position
}

type buttonPushes struct {
	a int
	b int
}

type Day13 struct {
	machines []clawMachine
	pushes   []buttonPushes
	line     string
	TenQ     bool
}

const aCost = 3
const bCost = 1
const aPre = "Button A: "
const bPre = "Button B: "
const priPre = "Prize: "

func (data *Day13) ParseDataLine(dataLine string) error {
	data.line = dataLine
	return nil
}

func (data *Day13) ProcessDataLine() error {
	if strings.HasPrefix(data.line, aPre) {
		var x, y int

		fmt.Sscanf(data.line[len(aPre):], "X+%d, Y+%d", &x, &y)
		data.machines = append(data.machines, clawMachine{a: Position{x, y}})
	}

	if strings.HasPrefix(data.line, bPre) {
		var x, y int

		fmt.Sscanf(data.line[len(bPre):], "X+%d, Y+%d", &x, &y)
		data.machines[len(data.machines)-1].b = Position{x, y}
	}

	if strings.HasPrefix(data.line, priPre) {
		var x, y int

		fmt.Sscanf(data.line[len(priPre):], "X=%d, Y=%d", &x, &y)
		data.machines[len(data.machines)-1].prize = Position{x, y}
	}

	return nil
}

func (data *Day13) ProcessDataSet() error {
	for _, machine := range data.machines {
		pushes := solveEquations(machine, data.TenQ)
		data.pushes = append(data.pushes, pushes)
	}
	return nil
}

func (data *Day13) Solve() (int, error) {
	total := 0
	for _, pushes := range data.pushes {
		total += ((pushes.a * aCost) + (pushes.b * bCost))
	}
	return total, nil
}

func solveEquations(machine clawMachine, tenQ bool) buttonPushes {
	var p1, p2 int64
	x1, y1, p1 := machine.a.x, machine.b.x, int64(machine.prize.x)
	x2, y2, p2 := machine.a.y, machine.b.y, int64(machine.prize.y)

	if tenQ {
		p1 += 10000000000000
		p2 += 10000000000000
	}

	// ignore lcm
	x1 *= y2
	p1 *= int64(y2)

	x2 *= y1
	p2 *= int64(y1)

	a := float64(p1-p2) / float64(x1-x2)
	if (!tenQ && a > 100) || a-float64(int(a)) > 0 {
		return buttonPushes{}
	}
	x2, y2, p2 = machine.a.y, machine.b.y, int64(machine.prize.y)
	if tenQ {
		p2 += 10000000000000
	}
	b := (float64(p2) - (a * float64(x2))) / float64(y2)
	if (!tenQ && b > 100) || b-float64(int(b)) > 0 {
		return buttonPushes{}
	}

	return buttonPushes{int(a), int(b)}
}

// 71x + 34y = 3685
// 86x + 11y = 2894

// 71x= 3685 - 34y
// x = 3685 - 24y
//     ---------
// 		    71

// 86(3685 - 34y) + 11y = 2894
//    ----------
// 	     71

// 316910 - 2924y + 11y = 2894
// ------   -----
//   71       71

// -2924y + 11y = 2894 - 316910
// ------                ------
//   71                    71

// -2924y + 781y = 205474 - 316910
// ------   ----   ------   ------
//   71      71      71       71

// -2924y + 781y = 205474 - 316910
// -2143y = -111436
// y = 52

// 71x + 1768 = 3685
// 71x = 1917
// x = 27

// (3685, 2894)

// mult largest by b button, mult smaller by a button, add results
