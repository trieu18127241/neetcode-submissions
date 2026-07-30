type Queue struct {
	element []int
}

func (this *Queue) Push(x int) {
	this.element = append(this.element, x)
}

func (this *Queue) Peek() int {
	if len(this.element) == 0 {
		return -1
	}
	return this.element[0]
}

func (this *Queue) Len() int {
	return len(this.element)
}

func dailyTemperatures(temperatures []int) []int {
	var result []int
	for i, v := range temperatures {
		var queue Queue
		var j int
		queue.Push(v)
		if i == len(temperatures)-1 {
			result = append(result, 0)
			break
		}
		for j = i + 1; j < len(temperatures) && queue.Peek() >= temperatures[j]; j++ {
			queue.Push(temperatures[j])
		}
		if j == len(temperatures) {
			result = append(result, 0)
		} else {
			result = append(result, queue.Len())
		}
	}
	return result
}