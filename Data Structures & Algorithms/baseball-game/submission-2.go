type Stack []int

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return 0
	}
	length := len(*s) - 1
	value := (*s)[length]
	*s = (*s)[:length]
	return value
}

func (s *Stack) Push(value int) {
	*s = append(*s, value)
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

func calPoints(operations []string) int {
	var stack Stack
	for _, v := range operations {
		if v == "+" {
			v1 := stack.Pop()
			v2 := stack.Pop()
			sum := v1 + v2
			stack.Push(v2)
			stack.Push(v1)
			stack.Push(sum)
			fmt.Println("+ case: ", stack)
		} else if v == "D" {
			v1 := stack.Pop()
			double := v1 * 2
			stack.Push(v1)
			stack.Push(double)
			fmt.Println("D case: ", stack)
		} else if v == "C" {
			stack.Pop()
			fmt.Println("C case: ", stack)
		} else {
			integerValue, _ := strconv.Atoi(v)
			stack.Push(integerValue)
		}
	}
	result := 0
	for _, v := range stack {
		result += v
	}
	return result
}