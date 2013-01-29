package alaska

import (
	"fmt"
)

type ColumnChunk struct {
	Dict     IntDictionary
	Elements []int
}

func (c *ColumnChunk) GetGlobalId(i int) int {
	val, err := c.Dict.ValueAt(c.Elements[i])
	if err != nil {
		panic(fmt.Errorf("alaska: ColumnChunk '%v' is corrupt", c))
	}
	return val
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
		ret[i] = c.GetGlobalId(idx)
	}
	return ret
}
