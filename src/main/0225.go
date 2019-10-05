package main

import "fmt"

func main() {
	var stack = Constructor()
	fmt.Println(stack.Empty())
	stack.Push(1)
	fmt.Println(stack.Top())
	fmt.Println(stack.Empty())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Empty())
}

type MyStack struct {
	size int
	nums []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{0, []int{}}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.size++
	this.nums = append(this.nums, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	var ret = this.nums[this.size-1]
	this.nums = this.nums[:this.size-1]
	this.size--
	return ret
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.nums[this.size-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.size == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
