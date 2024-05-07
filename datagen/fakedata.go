package datagen

import (
	"encoding/csv"
	"os"
)

func ReadCSV(filepath string, ch chan string) error {
	defer close(ch)
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err != nil {
			break
		}
		ch <- record[0]
	}

	return nil
}
