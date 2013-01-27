package main

type ColumnChunk struct {
	Dict IntDictionary
	Elements []int
}

func (c *ColumnChunk) GetGlobalId(i int) int {
	return c.Dict.ValueAt(c.Elements[i])
}
