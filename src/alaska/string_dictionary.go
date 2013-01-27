package main

type StringDictionary struct {
	Data []string
}

func (d *StringDictionary) IndexOf(v string) int {
	for i := 0; i < len(d.Data); i++ {
		if d.Data[i] == v {
			return i
		}
	}
	return -1
}

func (d *StringDictionary) ValueAt(i int) string {
	return d.Data[i]
}
