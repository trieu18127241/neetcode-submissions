func hasDuplicate(nums []int) bool {
    exist := make(map[int]bool)
    for i := 0; i < len(nums); i++ {
        fmt.Println("processing ", nums[i])
        _, ok := exist[nums[i]]; if !ok {
            fmt.Println("add num ", nums[i])
            exist[nums[i]] = true
        } else {
            fmt.Println("return ", nums[i])
            return true
        }
    }
    return false
}
