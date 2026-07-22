func calPoints(operations []string) int {
	stack := make([]int, 0)
	for _, v := range operations {
		if v == "+" {
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		} else if v == "D" {
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-1])
		} else if v == "C" {
			stack = stack[:len(stack)-1]
		} else {
			integerValue, _ := strconv.Atoi(v)
			stack = append(stack, integerValue)
		}
	}
	result := 0
	for _, v := range stack {
		result += v
	}
	return result
}