package Easy

/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.
Example:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
*/

type MinStack struct {
	nums []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{nums: make([]int, 0)}
}


func (this *MinStack) Push(x int)  {
	this.nums = append(this.nums, x)
}


func (this *MinStack) Pop()  {
	this.nums = this.nums[0:len(this.nums)-1]
}


func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}


func (this *MinStack) GetMin() int {
	min := this.nums[0]
	for i:=1;i<len(this.nums);i++{
		if min > this.nums[i]{
			min = this.nums[i]
		}
	}
	return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */