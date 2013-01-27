package main

type StringDictionary struct {
	data []string
}

func (d *StringDictionary) IndexOf(v string) int {
	for i := 0; i < len(d.data); i++ {
		if d.data[i] == v {
			return i
		}
	}
	return -1
}

func (d *StringDictionary) ValueAt(i int) string {
	return d.data[i]
}
