func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		num := nums[mid]

		if target == num {
			return mid
		} else if num > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}