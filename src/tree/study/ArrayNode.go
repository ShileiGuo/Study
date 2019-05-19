package study

import "fmt"

//ArrayStack ...
type ArrayStack struct {
	data []interface{}
	//栈顶指针
	top int
}

//NewArrayStack for initial
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

//IsEmpty for check if stack is empty
func (stack *ArrayStack) IsEmpty() bool {
	if stack.top < 0 || stack == nil {
		return true
	}
	return false
}

//Push for push data
func (stack *ArrayStack) Push(v interface{}) {
	if stack == nil {
		return
	}
	if stack.top < 0 {
		stack.top = 0
	} else {
		stack.top++
	}
	if stack.top > len(stack.data)-1 {
		stack.data = append(stack.data, v)
	} else {
		stack.data[stack.top] = v
	}
}

//Pop for read data
func (stack *ArrayStack) Pop() interface{} {
	if stack.IsEmpty() || stack == nil {
		return nil
	}
	v := stack.data[stack.top]
	stack.top--
	return v
}

//Top for read top data
func (stack *ArrayStack) Top() interface{} {
	if stack.IsEmpty() || stack == nil {
		return nil
	}
	return stack.data[stack.top]
}

//Flush is for clear data
func (stack *ArrayStack) Flush() {
	if stack.IsEmpty() || stack == nil {
		return
	}
	stack.top = -1
	stack.data = make([]interface{}, 0, 32)
	fmt.Println("Flush success")
}

//Print ...
func (stack *ArrayStack) Print() {
	if stack.IsEmpty() || stack == nil {
		return
	}
	for i := 0; i <= stack.top; i++ {
		fmt.Printf("The %d data is %v\n", i, stack.data[i])
	}
}

//PreOrderTraverseByArray ...
func (stack *ArrayStack) PreOrderTraverseByArray(rootIndex int) {
	if rootIndex > stack.top {
		return
	}

	fmt.Print(stack.data[rootIndex], " ")
	stack.PreOrderTraverseByArray(rootIndex*2 + 1)
	stack.PreOrderTraverseByArray(rootIndex*2 + 2)
}

//InOrderTraverseByArray ...
func (stack *ArrayStack) InOrderTraverseByArray(rootIndex int) {
	if rootIndex > stack.top {
		return
	}
	stack.InOrderTraverseByArray(rootIndex*2 + 1)
	fmt.Print(stack.data[rootIndex], " ")
	stack.InOrderTraverseByArray(rootIndex*2 + 2)
}

//PostOrderTraverseByArray ...
func (stack *ArrayStack) PostOrderTraverseByArray(rootIndex int) {
	if rootIndex > stack.top {
		return
	}
	stack.PostOrderTraverseByArray(rootIndex*2 + 1)
	stack.PostOrderTraverseByArray(rootIndex*2 + 2)
	fmt.Print(stack.data[rootIndex], " ")
}
