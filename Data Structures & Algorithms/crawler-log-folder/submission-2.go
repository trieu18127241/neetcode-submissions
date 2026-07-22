func minOperations(logs []string) int {
	deep := 0
	for _, action := range logs {
		step := actionDetect(action)
		fmt.Printf("action: %v, step: %v, deep: %v", action, step, deep)
		deep += step
		if deep < 0 {
			deep = 0
		}
		fmt.Printf("action: %v, deep: %v", action, deep)
	}
	return deep
}

func actionDetect(action string) int {
	if action == "../" {
		return -1
	}
	if action == "./" {
		return 0
	}
	return 1
}