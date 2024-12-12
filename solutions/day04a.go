package solutions

type Day04a struct {
	Lines []string
	Count int
}

var mas string = "MAS"

func (data *Day04a) ParseDataLine(dataLine string) error {
	data.Lines = append(data.Lines, dataLine)
	return nil
}

func (data *Day04a) ProcessDataLine() error {

	return nil
}

func (data *Day04a) ProcessDataSet() error {
	for y, line := range data.Lines {
		for x, chr := range line {
			if chr == 'X' {
				if checkForXmas(data.Lines, x, y, func(x, y int) (int, int) {
					return x, y - 1
				}) {
					data.Count++
				}
				if checkForXmas(data.Lines, x, y, func(x, y int) (int, int) {
					return x + 1, y - 1
				}) {
					data.Count++
				}
				if checkForXmas(data.Lines, x, y, func(x, y int) (int, int) {
					return x + 1, y
				}) {
					data.Count++
				}
				if checkForXmas(data.Lines, x, y, func(x, y int) (int, int) {
					return x + 1, y + 1
				}) {
					data.Count++
				}
				if checkForXmas(data.Lines, x, y, func(x, y int) (int, int) {
					return x, y + 1
				}) {
					data.Count++
				}
				if checkForXmas(data.Lines, x, y, func(x, y int) (int, int) {
					return x - 1, y + 1
				}) {
					data.Count++
				}
				if checkForXmas(data.Lines, x, y, func(x, y int) (int, int) {
					return x - 1, y
				}) {
					data.Count++
				}
				if checkForXmas(data.Lines, x, y, func(x, y int) (int, int) {
					return x - 1, y - 1
				}) {
					data.Count++
				}
			}
		}
	}
	return nil
}

func (data *Day04a) Solve() (int, error) {
	return data.Count, nil
}

func checkForXmas(lines []string, x, y int, adjust func(int, int) (int, int)) bool {
	for _, chr := range mas {
		x, y = adjust(x, y)
		if y < 0 || y >= len(lines) || x < 0 || x >= len(lines[y]) || rune(lines[y][x]) != chr {
			return false
		}
	}
	return true
}
