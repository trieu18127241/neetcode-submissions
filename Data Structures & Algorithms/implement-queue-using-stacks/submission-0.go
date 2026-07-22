type MyQueue struct {
	element []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.element = append(this.element, x)
}

func (this *MyQueue) Pop() int {
	if len(this.element) == 0 {
		return 0
	}
	firstElement := this.element[0]
	this.element = this.element[1:]
	return firstElement
}

func (this *MyQueue) Peek() int {
	if len(this.element) == 0 {
		return 0
	}
	return this.element[0]
}

func (this *MyQueue) Empty() bool {
	if len(this.element) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */