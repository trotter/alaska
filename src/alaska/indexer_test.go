package alaska

import (
	"testing"
	"strings"
)

func TestIndexingSimpleTable(t *testing.T) {
	indexer := GenerateIndexer()
	table := indexer.Run()

	t.Errorf("Table: %v", table)
}

func GenerateIndexer() Indexer {
	spec := NewIndexerSpec([]string{"State", "City"}, []string{"string","string"})
	rawData := "GA,Atlanta\nGA,Athens\nCA,San Francisco"
	return Indexer{"CityStates", spec, NewCsvSerde(), strings.NewReader(rawData)}
}
