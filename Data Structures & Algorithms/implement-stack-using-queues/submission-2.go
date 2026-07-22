type MyStack struct {
	element []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.element = append(this.element, x)
}

func (this *MyStack) Pop() int {
	fmt.Println("before pop: ", this.element)
	if len(this.element) == 0 {
		return 0
	}
	lastElement := this.element[len(this.element)-1]
	fmt.Println("info pop: ", lastElement)
	this.element = this.element[:len(this.element)-1]
	fmt.Println("after pop: ", this.element)
	return lastElement
}

func (this *MyStack) Top() int {
	if len(this.element) == 0 {
		return 0
	}
	return this.element[len(this.element)-1]
}

func (this *MyStack) Empty() bool {
	fmt.Println("empty: ", this.element)
	if len(this.element) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
