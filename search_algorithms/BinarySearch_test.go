package search_algorithms

import "testing"

func TestBinarySearchFindsTarget(t *testing.T) {
	target := 8
	nums := []int{1, 3, 5, 7, 8, 12, 13}
	bs := NewBinarySearch(nums)

	if !bs.LookUp(target) {
		t.Errorf("Expected %v to be found, but %v was not found instead\n", target, target)
	}
}
