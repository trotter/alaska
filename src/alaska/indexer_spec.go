package alaska

type IndexerSpec struct {
	ColumnNames []string
	ColumnTypes []string
}

func NewIndexerSpec(columnNames []string, columnTypes []string) *IndexerSpec {
	return &IndexerSpec{columnNames,columnTypes}
}
