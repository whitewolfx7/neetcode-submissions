type MinStack struct {
	stack []int
	minStack []int
}

func Constructor() MinStack {
	var minStack MinStack = MinStack{
		stack: []int{},
		minStack: []int{},
	}
	return minStack
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 || this.minStack[len(this.stack)-1] > val{
		this.stack = append(this.stack, val)
		this.minStack = append(this.minStack, val)
	}else{
		this.stack = append(this.stack, val)
		val = this.minStack[len(this.minStack)-1]
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack) - 1]
	this.minStack = this.minStack[:len(this.minStack) - 1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.stack)-1]
}