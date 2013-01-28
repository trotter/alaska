package main

type StringColumn struct {
	Dict StringDictionary
	Chunks []ColumnChunk
}

func (c *StringColumn) Where(s string) [][]int {
	index := c.Dict.IndexOf(s)
	ret := make([][]int, len(c.Chunks))

	for i, chunk := range c.Chunks {
		ret[i] = chunk.Where(index)
	}
	return ret
}
