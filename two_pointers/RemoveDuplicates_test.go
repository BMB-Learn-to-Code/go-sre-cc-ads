package two_pointers

import (
	"fmt"
	"slices"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2, 2, 3, 3, 4}

	RemoveDuplicates(nums)
	want := []int{1, 2, 3, 4}

	if !slices.Equal(nums[:len(want)], want) {
		t.Errorf("Expected {%v} to equals {%v}...", fmt.Sprint(nums[:len(want)]), fmt.Sprint(want))
	}
}
