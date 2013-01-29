package alaska

import (
	"fmt"
)

type IntDictionary struct {
	Data []int
}

func (d *IntDictionary) IndexOf(v int) int {
	for low, high := 0, len(d.Data) - 1; low < high; {
		mid := (low + high) / 2
		switch {
		case d.Data[mid] > v: high = mid - 1
		case d.Data[mid] < v: low = mid + 1
		case d.Data[mid] == v: return mid
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
