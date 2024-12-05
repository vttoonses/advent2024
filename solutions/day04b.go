package solutions

type Day04b struct {
	Lines []string
	Count int
}

func (data *Day04b) ParseDataLine(dataLine string) error {
	data.Lines = append(data.Lines, dataLine)
	return nil
}

func (data *Day04b) ProcessDataLine() error {

	return nil
}

func (data *Day04b) ProcessDataSet() error {
	for y, line := range data.Lines {
		for x, chr := range line {
			if chr == 'A' && checkForXdMases(data.Lines, x, y) {
				data.Count++
			}
		}
	}
	return nil
}

func (data *Day04b) Solve() (int, error) {
	return data.Count, nil
}

func checkForXdMases(lines []string, x, y int) bool {
	if x == 0 || y == 0 || x == len(lines[y])-1 || y == len(lines)-1 {
		return false
	}

	return checkForMas(lines, x-1, y-1, x+1, y+1) && checkForMas(lines, x+1, y-1, x-1, y+1)
}

func checkForMas(lines []string, x1, y1, x2, y2 int) bool {
	if lines[y1][x1] == 'M' && lines[y2][x2] == 'S' {
		return true
	}
	if lines[y1][x1] == 'S' && lines[y2][x2] == 'M' {
		return true
	}

	return false
}
