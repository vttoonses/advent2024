package solutions

import "fmt"

type robot struct {
	position Position
	velocity Position
}

type Day14 struct {
	robots  []robot
	Width   int
	Height  int
	Seconds int
	Tree    bool
}

func (data *Day14) ParseDataLine(dataLine string) error {
	var px, py, vx, vy int

	fmt.Sscanf(dataLine, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
	data.robots = append(data.robots, robot{position: Position{px, py}, velocity: Position{vx, vy}})

	return nil
}

func (data *Day14) ProcessDataLine() error {

	return nil
}

func (data *Day14) ProcessDataSet() error {
	if data.Tree {
		data.Seconds = 0
		for ; !data.isTree(); data.Seconds++ {
			for lcv := 0; lcv < len(data.robots); lcv++ {
				data.robots[lcv].move(data.Width, data.Height)
			}
		}
	} else {
		for ; data.Seconds > 0; data.Seconds-- {
			for lcv := 0; lcv < len(data.robots); lcv++ {
				data.robots[lcv].move(data.Width, data.Height)
			}
		}
	}
	return nil
}

func (data *Day14) Solve() (int, error) {
	if data.Tree {
		for y := 0; y < data.Height; y++ {
			for x := 0; x < data.Width; x++ {
				count := 0
				for _, r := range data.robots {
					if r.position.x == x && r.position.y == y {
						count += 1
					}
				}
				fmt.Print(count)
			}
			fmt.Println()
		}
		return data.Seconds, nil
	}

	q1, q2, q3, q4 := 0, 0, 0, 0
	w := int(data.Width / 2)
	h := int(data.Height / 2)

	for _, robot := range data.robots {
		if robot.position.x < w && robot.position.y < h {
			q1++
		}
		if robot.position.x < w && robot.position.y > h {
			q3++
		}
		if robot.position.x > w && robot.position.y < h {
			q2++
		}
		if robot.position.x > w && robot.position.y > h {
			q4++
		}
	}
	return q1 * q2 * q3 * q4, nil
}

func (r *robot) move(width, height int) {
	r.position.x += r.velocity.x
	r.position.y += r.velocity.y
	if r.position.x >= width {
		r.position.x -= width
	}
	if r.position.x < 0 {
		r.position.x = width + r.position.x
	}
	if r.position.y >= height {
		r.position.y -= height
	}
	if r.position.y < 0 {
		r.position.y = height + r.position.y
	}
}

func (data *Day14) isTree() bool {
	majority := int(float64(len(data.robots)) * 0.80)
	count := 0

	for lcv, r1 := range data.robots {
		for _, r2 := range data.robots[lcv + 1:] {
			if abs(r1.position.x - r2.position.x) <= 1 && abs(r1.position.y - r2.position.y) <= 1 {
				count += 1
			}
		}
	}

	return count >= majority
}
