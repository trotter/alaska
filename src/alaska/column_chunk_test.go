package alaska

import (
	"math/rand"
	"testing"
	"time"
)

func TestSelectReturnsGlobalIds(t *testing.T) {
	dict     := IntDictionary{[]int{15, 25, 10, 5}}
	elements := []int{3, 3, 1, 1, 2, 2, 0, 0}
	chunk    := ColumnChunk{dict, elements}

	expected := []int{5, 5, 15}
	actual   := chunk.Select([]int{0,1,7})

	if len(actual) != 3 {
		t.Errorf("Expected %v, got %v", 3, len(actual))
	}

	for i, el := range actual {
		if expected[i] != el {
			t.Errorf("Expected %v, got %v", expected[i], el)
		}
	}
}

// Unused, but keeping this around for now as a reference
// for generating random ColumnChunks
func GenerateColumnChunk(dictSize int, size int) ColumnChunk {
	rand := rand.New(rand.NewSource(time.Now().Unix()))
	dict := IntDictionary{rand.Perm(dictSize)}
	elements := make([]int, size)
	for i := 0; i < size; i++ {
		elements[i] = rand.Intn(dictSize)
	}
	return ColumnChunk{dict, elements}
}
