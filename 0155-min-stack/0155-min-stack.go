type MinStack struct {
    items []int
    mins []int
}


func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int)  {
    this.items = append(this.items, val)                // push to stack 
    if len(this.mins) == 0 || val <= this.GetMin() {    // if stack empty or new min, push new min
        this.mins = append(this.mins, val)
    }
}


func (this *MinStack) Pop()  {
    if len(this.items) == 0 { return }                  // skip for empty stack
    top := this.items[len(this.items)-1]                // var for top of stack
    this.items = this.items[:len(this.items)-1]         // pop top from stack
    if top == this.mins[len(this.mins)-1]{              // if popped was min, pop from min stack too
        this.mins = this.mins[:len(this.mins)-1]
    }               
}


func (this *MinStack) Top() int {
    if len(this.items) != 0 {
        return this.items[len(this.items)-1]             // if stack not empty, peek top item
    }
    return 0
}


func (this *MinStack) GetMin() int {
    if len(this.mins) != 0 {
        return this.mins[len(this.mins)-1]               // if stack not empty, peek curr min
    }
    return 0
    
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */