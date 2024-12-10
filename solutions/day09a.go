package solutions

type file struct {
	size   int
	blocks []int // -1 is free space
}

type Day09a struct {
	line     string
	storage  []file
	lastFile int
	ByFile   bool
}

func (data *Day09a) ParseDataLine(dataLine string) error {
	data.line = dataLine
	return nil
}

func (data *Day09a) ProcessDataLine() error {
	id := 0
	for lcv, digit := range data.line {
		size := int(digit - '0')

		var blocks []int = make([]int, size)

		val := -1
		if lcv%2 == 0 {
			data.lastFile = lcv
			val = id
			id++
		}

		for lcv := 0; lcv < size; lcv++ {
			blocks[lcv] = val
		}

		data.storage = append(data.storage, file{
			size:   size,
			blocks: blocks,
		})
	}

	return nil
}

func (data *Day09a) ProcessDataSet() error {
	if data.ByFile {
		for {
			if data.lastFile <= 1 {
				return nil
			}

			firstSpace := data.findSpace()
			if firstSpace < data.lastFile {
				data.moveFile(firstSpace)
			}
			data.lastFile -= 2
		}
	} else {
		firstSpace := 1
		for {
			if firstSpace > data.lastFile {
				return nil
			}

			firstSpace = data.moveData(firstSpace)
			data.lastFile -= 2
		}
	}
}

func (data *Day09a) Solve() (int, error) {
	checkSum := 0
	idx := 0

	for _, dataFile := range data.storage {
		for _, block := range dataFile.blocks {
			if data.ByFile {
				if block > -1 {
					checkSum += (idx * block)
				}
			} else {
				if block == -1 {
					return checkSum, nil
				}

				checkSum += (idx * block)
			}
			idx++
		}
	}

	return checkSum, nil
}

func (data *Day09a) moveData(firstSpace int) int {
	dIdx := len(data.storage[data.lastFile].blocks) - 1

	for ; dIdx >= 0; dIdx-- {
		fIdx := nextFree(data.storage[firstSpace].blocks)
		for {
			if fIdx < len(data.storage[firstSpace].blocks) {
				break
			}

			firstSpace += 2
			if firstSpace > data.lastFile {
				return firstSpace
			}
			fIdx = nextFree(data.storage[firstSpace].blocks)
		}

		data.storage[firstSpace].blocks[fIdx] = data.storage[data.lastFile].blocks[dIdx]
		data.storage[data.lastFile].blocks[dIdx] = -1
	}

	return firstSpace
}

func nextFree(blocks []int) int {
	for lcv, val := range blocks {
		if val == -1 {
			return lcv
		}
	}
	return len(blocks)
}

func (data *Day09a) findSpace() int {
	dataLen := len(data.storage[data.lastFile].blocks)
	for lcv := 1; lcv < data.lastFile; lcv += 2 {
		availableSpace := 0
		for _, val := range data.storage[lcv].blocks {
			if val == -1 {
				availableSpace += 1
			}
		}
		if availableSpace >= dataLen {
			return lcv
		}
	}
	return data.lastFile + 1
}

func (data *Day09a) moveFile(firstSpace int) {
	space := 0
	for lcv, val := range data.storage[firstSpace].blocks {
		if val == -1 {
			space = lcv
			break
		}
	}
	for lcv, val := range data.storage[data.lastFile].blocks {
		data.storage[firstSpace].blocks[space+lcv] = val
		data.storage[data.lastFile].blocks[lcv] = -1
	}
}
