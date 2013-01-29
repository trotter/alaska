package alaska

import (
	"fmt"
)

type Table struct {
	Name    string
	Columns []StringColumn
}

func (c *Table) Select(indices [][]int) [][][]string {
	ret := make([][][]string, len(c.Columns))
	for i, column := range c.Columns {
		ret[i] = column.Select(indices)
	}
	return ret
}

func (c *Table) Where(columnName string, term string) ([][]int, error) {
	column, err := c.findColumn(columnName)
	if err != nil {
		return [][]int{}, err
	}
	return column.Where(term), nil
}

func (c *Table) findColumn(columnName string) (StringColumn, error) {
	for _, column := range c.Columns {
		if columnName == column.Name {
			return column, nil
		}
	}
	return StringColumn{}, fmt.Errorf("alaska: Could not find column '%s' in table '%s'", columnName, c.Name)
}
