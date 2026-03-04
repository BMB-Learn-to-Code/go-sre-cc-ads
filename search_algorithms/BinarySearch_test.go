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

// Utility Funciont to check the method bellow
func checkIfOdd(target int) bool {
	// Todo Condition
	return target%2 > 0

}
func TestBinarySearchFindsBoudary(t *testing.T) {
	nums := []int{1, 13, 9, 7, 100, 2, 5, 6}
	want := 3

	bs := NewBinarySearch(nums)

	if idx := bs.FindBoundary(checkIfOdd); idx != want {
		t.Errorf("Expected boundary to be found at index: %v, but found %v ", want, idx)
	}

}
