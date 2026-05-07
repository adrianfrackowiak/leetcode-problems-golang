func containsDuplicate(nums []int) bool {
	m := make(map[int]int)

	for _, n := range nums {
		m[n]++
		if m[n] > 1 {
			return true
		}
	}

	return false
}