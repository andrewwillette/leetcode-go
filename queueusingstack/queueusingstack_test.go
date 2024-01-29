package queueusingstack

// https://leetcode.com/problems/implement-queue-using-stacks
type MyQueue struct {
    stackA []int
}

func Constructor() MyQueue {
    return MyQueue{stackA: []int{}, stackB: []int{}}
}

func (this *MyQueue) Push(x int)  {
    this.stackA = append(this.stackA, x)
}

func (this *MyQueue) Pop() int {
    stackAUpdated := make([]int, len(this.stackA)-1)
    for i :=1;i<len(this.stackA); i++ {
        stackAUpdated[i-1] = this.stackA[i]
    }
    toPop := this.stackA[0]
    this.stackA = stackAUpdated
    return toPop
}

func (this *MyQueue) Peek() int {
    return this.stackA[0]
}

func (this *MyQueue) Empty() bool {
    return len(this.stackA) == 0
}

