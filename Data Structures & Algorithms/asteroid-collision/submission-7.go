type Stack struct {
	element []int
}

func Constructor(list []int) Stack {
	return Stack{
		element: list,
	}
}

func (this *Stack) Push(x int) {
	this.element = append(this.element, x)
}

func (this *Stack) Pop() int {
	if len(this.element) == 0 {
		return -1
	}
	lastElement := this.element[len(this.element)-1]
	this.element = this.element[:len(this.element)-1]
	return lastElement
}

func (this *Stack) Peek() int {
	if len(this.element) == 0 {
		return -1
	}
	return this.element[len(this.element)-1]
}

func asteroidCollision(asteroids []int) []int {
	var stack Stack
	for _, v := range asteroids {
		if stack.Peek()*v < 0 {
            if stack.Peek() < 0 {
				stack.Push(v)
			} else {
                var flag bool
                for stack.Peek()*v < 0 {
                    if math.Abs(float64(stack.Peek())) < math.Abs(float64(v)) && stack.Peek() >= 0 {
                        stack.Pop()
                    } else {
                        flag = true
                        if math.Abs(float64(stack.Peek())) == math.Abs(float64(v)) {
                            stack.Pop()
                        }
                        break
                    }
                }
                if !flag {
                    stack.Push(v)
                }
            }
		} else {
			stack.Push(v)
		}
	}
	return stack.element
}