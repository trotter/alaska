package alaska

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

func (c *StringColumn) Select(indices [][]int) [][]string {
	ret := make([][]string, len(c.Chunks))
	for i, chunk := range c.Chunks {
		globalIds := chunk.Select(indices[i])
		ret[i] = make([]string, len(globalIds))
		for j, id := range globalIds {
		  ret[i][j] = c.Dict.ValueAt(id)
		}
	}
	return ret
}
