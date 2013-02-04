package alaska

import (
	"io"
)

type Indexer struct {
	Name        string
	Spec        *IndexerSpec
	Serde       Serde
	RawData     io.Reader
}

func (i *Indexer) Run() Table {
	return Table{"blah", []StringColumn{}}
}
