package alaska

import (
	"fmt"
)

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

func (c *StringColumn) Select(indices [][]int) ([]string, error) {
	ret := []string{}
	for i, chunk := range c.Chunks {
		globalIds, err := chunk.Select(indices[i])
		if err != nil {
			return []string{}, err
		}
		for _, id := range globalIds {
			val, err := c.Dict.ValueAt(id)
			if err != nil {
				panic(fmt.Errorf("alaska: Column '%v' is corrupt", c))
			}
			ret = append(ret, val)
		}
	}
	return ret, nil
}
