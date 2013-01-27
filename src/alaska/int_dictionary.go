package main

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

func (d *IntDictionary) ValueAt(i int) int {
	return d.Data[i]
}
