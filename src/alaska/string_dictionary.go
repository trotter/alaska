package alaska

import(
	"fmt"
)

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

func (d *StringDictionary) ValueAt(i int) (string, error) {
	if i > len(d.Data) {
		return "", fmt.Errorf("stringDictionary: index out of bounds - i=%d,len=%d", i, len(d.Data))
	}
  return d.Data[i], nil
}
