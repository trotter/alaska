package alaska

type StringColumn struct {
	Name   string
	Dict   StringDictionary
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

func (c *StringColumn) Select(indices [][]int) []string {
	ret := []string{}
	for i, chunk := range c.Chunks {
		globalIds := chunk.Select(indices[i])
		for _, id := range globalIds {
			val, _ := c.Dict.ValueAt(id)
			ret = append(ret, val)
		}
	}
	return ret
}
