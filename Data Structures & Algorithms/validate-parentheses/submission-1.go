type Stack struct {
	top  int
	data []string
}

func push(stack *Stack, data string) {
	stack.top++
	stack.data = append(stack.data, data)
}

func pop(stack *Stack) {
	if len(stack.data) == 0 {
		return
	}
	stack.top--
	last := stack.data[stack.top]
	fmt.Println("last:", last)
	stack.data = stack.data[:stack.top]
}

func peek(stack *Stack) string {
    if len(stack.data) == 0 {
        return ""
    }
	return stack.data[stack.top-1]
}

func empty(stack *Stack) bool {
	return stack.top == 0
}

func isValid(s string) bool {
	stack := Stack{}
	mapClosure := map[string]string{"}": "{", ")": "(", "]": "["}
	for _, char := range s {
		val, exist := mapClosure[string(char)]
		if exist {
            fmt.Println("exist:", val)
			peekData := peek(&stack)
			if peekData != val {
				return false
			} else {
				pop(&stack)
			}
		} else {
			push(&stack, string(char))
		}
	}
	return empty(&stack)
}