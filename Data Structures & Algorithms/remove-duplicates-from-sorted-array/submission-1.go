func removeDuplicates(nums []int) int {
	l := 0
	for l < len(nums)-1 {
		r := l + 1
		if nums[l] == nums[r] {
			fmt.Println(l, r, nums)
			count := 1
			for r+count < len(nums) && nums[r] == nums[r+count] {
				count++
			}
			fmt.Println(l, r, nums, count)
			for r+count < len(nums) {
				nums[r] = nums[r+count]
				r++
			}
			nums = nums[:len(nums)-count]
			fmt.Println(l, nums)
		}
		l++
	}
	return len(nums)
}