type MinStack struct {
	top  int
	data []int
}

func Constructor() MinStack {
	return MinStack{
		top:  0,
		data: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.top++
	this.data = append(this.data, val)
}

func (this *MinStack) Pop() {
	if len(this.data) > 0 {
		this.top--
		this.data = this.data[:this.top]
	}
}

func (this *MinStack) Top() int {
	if len(this.data) > 0 {
		return this.data[this.top-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	min := this.data[this.top-1]
	for i := len(this.data) - 1; i >= 0; i-- {
		if this.data[i] < min {
			min = this.data[i]
		}
	}
	return min
}
