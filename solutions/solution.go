package solutions

type Solution interface {
	ParseDataLine(string) error
	ProcessDataLine() error
	ProcessDataSet() error
	Solve() (int, error)
}

// func (data *Day01a) ParseDataLine(dataLine string) error {

// 	return nil
// }

// func (data *Day01a) ProcessDataLine() error {

// 	return nil
// }

// func (data *Day01a) ProcessDataSet() error {

// 	return nil
// }

// func (data *Day01a) Solve() (int, error) {

// 	return 0, nil
// }
