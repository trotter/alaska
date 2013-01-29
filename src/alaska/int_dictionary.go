package alaska

import (
	"fmt"
)

type IntDictionary struct {
	Data []int
}

func (d *IntDictionary) IndexOf(v int) int {
	for i := 0; i < len(d.Data); i++ {
		if d.Data[i] == v {
			return i
		}
	}
	return -1
}

func (d *IntDictionary) ValueAt(i int) (int, error) {
	if i > len(d.Data) {
		return 0, fmt.Errorf("intDictionary: index out of bounds - i=%d,len=%d", i, len(d.Data))
	}
	return d.Data[i], nil
}
