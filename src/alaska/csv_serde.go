package alaska

import (
	"encoding/csv"
	"strings"
)

type CsvSerde struct {
}

func NewCsvSerde() *CsvSerde {
	return &CsvSerde{}
}

func (s *CsvSerde) ProcessLine(line string) []string {
	// Terribly inefficient, but it works
	reader := csv.NewReader(strings.NewReader(line))

	// TODO: stop ignoring this error
	record, _ := reader.Read()
	return record
}
