package main

type Dictionary struct {
	data []string
}

func (d *Dictionary) IndexOf(v string) int {
	for i := 0; i < len(d.data); i++ {
		if d.data[i] == v {
			return i
		}
	}
	return -1
}

func (d *Dictionary) ValueAt(i int) string {
	return d.data[i]
}
