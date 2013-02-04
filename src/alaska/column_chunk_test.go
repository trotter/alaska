package alaska

import (
	"testing"
)

func TestSelectReturnsGlobalIds(t *testing.T) {
	chunk    := GenerateColumnChunk()
	expected := []int{5, 5, 15}
	actual   := chunk.Select([]int{0,1,7})

	if len(actual) != 3 {
		t.Errorf("Expected length to be %v, got %v", 3, len(actual))
	}

	for i, el := range actual {
		if expected[i] != el {
			t.Errorf("Expected element at %d to be %v, got %v", i, expected[i], el)
		}
	}
}

func TestWhereReturnsElementPositions(t *testing.T) {
	chunk    := GenerateColumnChunk()
	expected := []int{2, 3}
	actual   := chunk.Where(25)

	if len(actual) != 2 {
		t.Errorf("Expected length to be %v, got %v", 2, len(actual))
	}

	for i, el := range actual {
		if expected[i] != el {
			t.Errorf("Expected element at %d to be %v, got %v", i, expected[i], el)
		}
	}
}

func GenerateColumnChunk() ColumnChunk {
	dict     := IntDictionary{[]int{15, 25, 10, 5}}
	elements := []int{3, 3, 1, 1, 2, 2, 0, 0}
	return ColumnChunk{dict, elements}
}
