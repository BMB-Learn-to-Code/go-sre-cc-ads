package two_pointers

func RemoveDuplicates(nums []int) {
	left := 0

	for right := 1; right < len(nums); right++ {
		if nums[right] != nums[left] {
			left += 1
			nums[right], nums[left] = nums[left], nums[right]
		}
	}
}
