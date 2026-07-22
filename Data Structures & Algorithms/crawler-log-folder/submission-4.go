func minOperations(logs []string) int {
	stack := []string{}
	for _, log := range logs {
		if log == "../"{
			fmt.Println("before: ", stack)
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			fmt.Println("after: ", stack)
		} else if log != "./" {
			stack = append(stack, log)
		}
	}
	return len(stack)
}
