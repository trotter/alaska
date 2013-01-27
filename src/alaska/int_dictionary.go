package main

type IntDictionary struct {
	data []int
}

func (d *IntDictionary) IndexOf(v int) int {
	for i := 0; i < len(d.data); i++ {
		if d.data[i] == v {
			return i
		}
	}
	return -1
}

func (d *IntDictionary) ValueAt(i int) int {
	return d.data[i]
}
