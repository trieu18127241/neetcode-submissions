type DoubleLinkedList struct {
	val  string
	next *DoubleLinkedList
	prev *DoubleLinkedList
}

func evalRPN(tokens []string) int {
	head := &DoubleLinkedList{
		val: tokens[0],
	}
	curr := head
	for i := 1; i < len(tokens); i++ {
		newNode := &DoubleLinkedList{
			val:  tokens[i],
			prev: curr,
		}
		curr.next = newNode
		curr = curr.next
	}
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			left, _ := strconv.Atoi(head.prev.prev.val)
			right, _ := strconv.Atoi(head.prev.val)
			var result int
			switch tokens[i] {
			case "+":
				result = left + right
			case "-":
				result = left - right
			case "*":
				result = left * right
			case "/":
				result = left / right
			}
			head.val = strconv.Itoa(result)
			head.prev = head.prev.prev.prev
			head = head.next
		} else {
			head = head.next
		}
	}
    res, _ := strconv.Atoi(curr.val)
	return res
}