package s0232

import "fmt"

func Run0232() {
	var queue = Constructor()
	fmt.Println(queue.Empty())
	queue.Push(1)
	fmt.Println(queue.Peek())
	fmt.Println(queue.Empty())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Empty())
}

type MyQueue struct {
	nums []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{[]int{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.nums = append(this.nums, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	var ret = this.nums[0]
	this.nums = this.nums[1:]
	return ret
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.nums[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.nums) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
