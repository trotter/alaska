package alaska

type ColumnChunk struct {
	Dict IntDictionary
	Elements []int
}

func (c *ColumnChunk) GetGlobalId(i int) (int, error) {
	return c.Dict.ValueAt(c.Elements[i])
}

func (c *ColumnChunk) Where(globalIndex int) []int {
	ret := []int{}
	localIndex := c.Dict.IndexOf(globalIndex)
	for i, el := range c.Elements {
		if el == localIndex {
			ret = append(ret, i)
		}
	}
	return ret
}

func (c *ColumnChunk) Select(indices []int) []int {
	ret := make([]int, len(indices))
	for i, idx := range indices {
		ret[i] = c.Elements[idx]
	}
	return ret
}
