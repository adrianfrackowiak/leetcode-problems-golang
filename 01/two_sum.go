func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

    for i, v := range nums {
        var rest int = target - v
        if _, ok := m[rest]; ok {
            return []int{i, m[rest]}
        }
        m[v] = i
    }

    return []int{}
}
