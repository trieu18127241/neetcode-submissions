func removeDuplicates(nums []int) int {
	l := 0
	for l < len(nums)-1 {
		r := l + 1
		if nums[l] == nums[r] {
			count := 1
			for r+count < len(nums) && nums[r] == nums[r+count] {
				count++
			}
			for r+count < len(nums) {
				nums[r] = nums[r+count]
				r++
			}
			nums = nums[:len(nums)-count]
		}
		l++
	}
	return len(nums)
}