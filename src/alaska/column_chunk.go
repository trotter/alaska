package main

type ColumnChunk struct {
	Dict IntDictionary
	Elements []int
}

func (c *ColumnChunk) GetGlobalId(i int) int {
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
