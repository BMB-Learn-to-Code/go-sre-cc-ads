package search_algorithms

type BinarySearch struct {
	data []int
}

func NewBinarySearch(data []int) BinarySearch {
	return BinarySearch{data: data}
}

func (b *BinarySearch) LookUp(target int) bool {
	_, found := b.search(target)
	return found
}

func (b *BinarySearch) FindBoundary(condition func(n int) bool) int {
	return b.searchLast(condition)
}

func (b *BinarySearch) search(target int) (int, bool) {
	low := 0
	high := len(b.data) - 1

	for low <= high {
		mid := (high-low)/2 + low

		if b.data[mid] == target {
			return mid, true
		}
		if b.data[mid] > target {
			high = mid - 1
		}
		if b.data[mid] < target {
			low = mid + 1
		}
	}

	return -1, false
}

func (b *BinarySearch) searchLast(f func(n int) bool) int {
	low := 0
	high := len(b.data) - 1
	lastfound := -1

	for low <= high {
		mid := (high-low)/2 + low
		if f(b.data[mid]) {
			lastfound = mid
			low = mid + 1
		}
		high = mid - 1
	}

	return lastfound

}
