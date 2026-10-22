func sortedSquares(nums []int) []int {
	l, r := 0, len(nums)-1
	result := make([]int, 0)
	for l <= r {
		if nums[l]*nums[l] < nums[r]*nums[r] {
			result = append(result, nums[r]*nums[r])
			r--
		} else {
			result = append(result, nums[l]*nums[l])
			l++
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}