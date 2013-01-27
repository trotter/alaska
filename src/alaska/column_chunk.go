package main

type ColumnChunk struct {
	dict IntDictionary
	elements []int
}

func (c *ColumnChunk) GetGlobalId(i int) int {
	return c.dict.ValueAt(c.elements[i])
}
