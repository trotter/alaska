package alaska

import (
	"fmt"
)

type ColumnChunk struct {
	Dict     IntDictionary
	Elements []int
}

func (c *ColumnChunk) GetGlobalId(i int) (int, error) {
	if i > len(c.Elements) {
		return -1, fmt.Errorf("alaska: index out of bounds - i=%d;len=%d", i, len(c.Elements))
	}
	val, err := c.Dict.ValueAt(c.Elements[i])
	if err != nil {
		panic(fmt.Errorf("alaska: ColumnChunk '%v' is corrupt", c))
	}
	return val, nil
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

func (c *ColumnChunk) Select(indices []int) ([]int, error) {
	ret := make([]int, len(indices))
	for i, idx := range indices {
		val, err := c.GetGlobalId(idx)
		if err != nil {
			return []int{}, fmt.Errorf("alaska: invalid index in Select - idx=%v", idx)
		}
		ret[i] = val
	}
	return ret, nil
}
