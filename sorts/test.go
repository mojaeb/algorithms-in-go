package sorts

import (
	"testing"
)

var array = []int{3, 52, 11, 43, 53, 2, 43, 23, 543, 23, 3, 6, 64}

func sample() int {
	return 3
}

func TestInsersion(t *testing.T) {
	sample()
	t.Run("test insertion sort", func(t *testing.T) {
		result := sample()
		if result == 3 {
			t.Errorf("most return 3 and result is %d", result)
		}
	})
}
